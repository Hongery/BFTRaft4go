package server

import (
	"context"
	spb "github.com/PomeloCloud/BFTRaft4go/proto/server"
	"github.com/PomeloCloud/BFTRaft4go/utils"
	"github.com/golang/protobuf/proto"
	"github.com/dgraph-io/badger"
	"log"
	"sync"
	"github.com/tevino/abool"
	"crypto/x509"
)

// Alpha group is a group specialized for tracking network members and groups
// All nodes on the network should observe alpha group to provide group routing
// Alpha group will not track leadership changes, only members
// It also responsible for group creation and limit number of members in each group
// Both clients and cluster nodes can benefit form alpha group by bootstrapping with any node in the cluster
// It also provide valuable information for consistent hashing and distributed hash table implementations

// This file contains all of the functions for cluster nodes to track changes in alpha group

func (s *BFTRaftServer) ColdStart() {
	// cloud start will assign the node as the only member in it's alpha group
	alphaGroup := &spb.RaftGroup{
		Id: utils.ALPHA_GROUP,
		Replications: 32,
		LeaderPeer: s.Id,
		Term: 0,
	}
	thisPeer := &spb.Peer{
		Id: s.Id,
		Group: utils.ALPHA_GROUP,
		Host: s.Id,
		NextIndex: 0,
		MatchIndex: 0,
	}
	thisHost := &spb.Host{
		Id: s.Id,
		LastSeen: 0,
		Online: true,
		ServerAddr: s.Opts.Address,
	}
	thisHost.PublicKey, _ = x509.MarshalPKIXPublicKey(utils.PublicKeyFromPrivate(s.PrivateKey))
	if err := s.DB.Update(func(txn *badger.Txn) error {
		if err := s.SaveGroup(txn, alphaGroup); err != nil {
			return err
		}
		if err := s.SavePeer(txn, thisPeer); err != nil {
			return err
		}
		return s.SaveHost(txn, thisHost)
	}); err != nil {
		log.Fatal("cannot save to cold start:", err)
	}
	s.GroupsOnboard[utils.ALPHA_GROUP] = &RTGroupMeta{
		Peer:       thisPeer.Id,
		Leader:     alphaGroup.LeaderPeer,
		Lock:       sync.RWMutex{},
		GroupPeers: map[uint64]*spb.Peer{thisPeer.Id: thisPeer},
		Group:      alphaGroup,
		IsBusy:     abool.NewBool(false),
	}
	s.Client.AlphaRPCs.ResetBootstrap([]string{s.Opts.Address})
}

func (s *BFTRaftServer) SyncAlphaGroup() {
	// Force a snapshot sync for group members by asking alpha nodes for it
	// This function should be invoked every time it startup
	// First we need to get all alpha nodes
	alphaRPCs := s.Client.AlphaRPCs
	// get alpha peers from alpha nodes
	res := utils.MajorityResponse(alphaRPCs.Get(), func(client spb.BFTRaftClient) (interface{}, []byte) {
		if res, err := client.GroupPeers(context.Background(), &spb.GroupId{
			GroupId: utils.ALPHA_GROUP,
		}); err == nil {
			return res, GetPeersSignData(res.Peers)
		} else {
			return nil, []byte{}
		}
	})
	var alphaPeersRes *spb.GroupPeersResponse = nil
	if res == nil {
		alphaPeersRes = nil
	} else {
		alphaPeersRes = res.(*spb.GroupPeersResponse)
	}
	if alphaPeersRes == nil {
		log.Println("cannot get alpha peers, will try to cold start")
		s.ColdStart()
		return
	}
	peers := alphaPeersRes.Peers
	isAlphaMember := false
	for _, p := range peers {
		if p.Host == s.Id {
			isAlphaMember = true
			break
		}
	}
	lastEntry := alphaPeersRes.LastEntry
	group := s.GetGroupNTXN(utils.ALPHA_GROUP)
	if isAlphaMember {
		if group == nil {
			panic("Alpha member cannot find alpha group")
		}
		// Nothing should be done here, the raft algorithm should take the rest
	} else {
		if group == nil {
			log.Println("cannot find alpha group at local, will pull from remote")
			// alpha group cannot be found, it need to be generated
			res := utils.MajorityResponse(alphaRPCs.Get(), func(client spb.BFTRaftClient) (interface{}, []byte) {
				if res, err := client.GetGroupContent(context.Background(), &spb.GroupId{GroupId: utils.ALPHA_GROUP}); err == nil {
					if data, err2 := proto.Marshal(res); err2 == nil {
						return res, data
					} else {
						return nil, []byte{}
					}
				} else {
					log.Println(err)
					return nil, []byte{}
				}
			})
			if res != nil {
				group = res.(*spb.RaftGroup)
				log.Println("pulled alpha group at term:", group.Term, "leader:", group.LeaderPeer)
			} else {
				log.Println("cannot get alpha group from cluster")
			}
		}
		if group != nil {
			var lastIndex uint64
			if lastEntry == nil {
				group.Term = 0
				lastIndex = 0
			} else {
				group.Term = lastEntry.Index
				lastIndex = lastEntry.Index
			}
			s.DB.Update(func(txn *badger.Txn) error {
				s.SetGroupLogLastIndex(txn, utils.ALPHA_GROUP, lastIndex)
				// the index will be used to observe changes
				s.SaveGroup(txn, group)
				for _, peer := range peers {
					s.SavePeer(txn, peer)
				}
				return nil
			})
		} else {
			log.Fatal("cannot generate alpha group from cluster")
		}
	}
}