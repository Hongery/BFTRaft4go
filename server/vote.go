package server

import (
	"context"
	"fmt"
	pb "github.com/PomeloCloud/BFTRaft4go/proto/server"
	"github.com/PomeloCloud/BFTRaft4go/utils"
	"github.com/dgraph-io/badger"
	"log"
	"sync"
	"time"
)

func RequestVoteRequestSignData(req *pb.RequestVoteRequest) []byte {
	return []byte(fmt.Sprint(req.Group, "-", req.Term, "-", req.LogIndex, "-", req.Term, "-", req.CandidateId))
}

func RequestVoteResponseSignData(res *pb.RequestVoteResponse) []byte {
	return []byte(fmt.Sprint(res.Group, "-", res.Term, "-", res.LogIndex, "-", res.Term, "-", res.CandidateId, "-", res.Granted))
}

func (m *RTGroup) ResetTerm(term uint64) {
	m.Group.Term = term
	m.Votes = []*pb.RequestVoteResponse{}
	m.VotedPeer = 0
	for peerId := range m.GroupPeers {
		m.SendVotesForPeers[peerId] = true
	}
	m.Server.SaveGroupNTXN(m.Group)
}

func (m *RTGroup) BecomeCandidate() {
	m.RefreshTimer(10)
	m.Role = CANDIDATE
	group := m.Group
	m.ResetTerm(group.Term + 1)
	term := group.Term
	m.Server.SaveGroupNTXN(m.Group)
	lastEntry := m.LastLogEntryNTXN()
	var lastIndex uint64 = 0
	var lastLogTerm uint64 = 0
	if lastEntry != nil {
		lastIndex = lastEntry.Index
		lastLogTerm = lastEntry.Term
	}
	request := &pb.RequestVoteRequest{
		Group:       group.Id,
		Term:        term,
		LogIndex:    lastIndex,
		LogTerm:     lastLogTerm,
		CandidateId: m.Server.Id,
		Signature:   []byte{},
	}
	log.Println("become a candidate", ", term", m.Group.Term)
	request.Signature = m.Server.Sign(RequestVoteRequestSignData(request))
	voteReceived := make(chan *pb.RequestVoteResponse)
	numPeers := len(m.GroupPeers)
	wg := sync.WaitGroup{}
	wg.Add(numPeers)
	log.Println("sending vote request to", numPeers, "peers")
	for _, peer := range m.GroupPeers {
		nodeId := peer.Id
		node := m.Server.GetHostNTXN(nodeId)
		if node == nil {
			log.Println("skip requesting self for a vote")
			wg.Done()
			continue
		}
		go func() {
			if client, err := utils.GetClusterRPC(node.ServerAddr); err == nil {
				if voteResponse, err := client.RequestVote(context.Background(), request); err == nil {
					publicKey := m.Server.GetHostPublicKey(nodeId)
					signData := RequestVoteResponseSignData(voteResponse)
					if err := utils.VerifySign(publicKey, voteResponse.Signature, signData); err == nil {
						if voteResponse.Granted && voteResponse.LogIndex <= lastEntry.Index {
							voteReceived <- voteResponse
						} else {
							log.Println(nodeId, "peer not granted vote")
						}
					} else {
						log.Println("error on verify vote response:", err)
					}
				} else {
					log.Println("error on request vote:", err)
				}
			} else {
				log.Println("cannot get client for request votes")
			}
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(voteReceived)
		log.Println("received all vote response")
	}()
	adequateVotes := make(chan bool, 1)
	go func() {
		// Here we can follow the rule of Raft by expecting majority votes
		// or follow the PBFT rule by expecting n - f votes
		// I will use the rule from Raft first
		expectedVotes := len(m.GroupPeers) / 2 // ExpectedHonestPeers(s.OnboardGroupPeersSlice(group.Id))
		votes := 0
		for vote := range voteReceived {
			votes++
			log.Println("already  received", votes, "votes")
			if votes >= expectedVotes {
				m.Votes = append(m.Votes, vote)
				adequateVotes <- true
				break
			}
		}
	}()
	select {
	case <-adequateVotes:
		log.Println("now transfer to leader, term", m.Group.Term)
		m.BecomeLeader()
	case <-time.After(10 * time.Second):
		log.Println("vote requesting time out")
	}
}

func (m *RTGroup) BecomeLeader() {
	// when this peer become the leader of the group
	// it need to send it's vote to followers to claim it's authority
	// this only need to be done once in each term
	// so we just send the 'AppendEntry' request in this function
	// we can use a dedicated rpc protocol for this, but no bother
	m.Role = LEADER
	m.Leader = m.Server.Id // set self to leader for next following requests
	m.Server.DB.Update(func(txn *badger.Txn) error {
		return m.Server.SaveGroup(txn, m.Group)
	})
	m.SendFollowersHeartbeat(context.Background())
}

func (m *RTGroup) BecomeFollower(appendEntryReq *pb.AppendEntriesRequest) bool {
	// first we need to verify the leader got all of the votes required
	log.Println("trying to become a follower of", appendEntryReq.LeaderId, ", term", appendEntryReq.Term)
	expectedVotes := ExpectedHonestPeers(m.OnboardGroupPeersSlice())
	receivedVotes := len(appendEntryReq.QuorumVotes)
	if receivedVotes < expectedVotes {
		log.Println("did not received enough vote", receivedVotes, "/", expectedVotes)
		return false
	}
	votes := map[uint64]bool{}
	for _, vote := range appendEntryReq.QuorumVotes {
		votePeer, foundCandidate := m.GroupPeers[vote.Voter]
		if !foundCandidate || vote.Term <= m.Group.Term {
			log.Println("invalid candidate:", vote.Voter)
			continue
		}
		// check their signatures
		signData := RequestVoteResponseSignData(vote)
		publicKey := m.Server.GetHostPublicKey(votePeer.Id)
		if err := utils.VerifySign(publicKey, vote.Signature, signData); err != nil {
			log.Println("verify vote from", vote.Voter, "failed:", err)
			continue
		}
		// check their properties to avoid forging
		if vote.Group == m.Group.Id && vote.CandidateId == appendEntryReq.LeaderId && vote.Granted {
			votes[votePeer.Id] = true
		} else {
			log.Println("vote properity not match this vote term, grant:", vote.Granted)
		}
	}
	if len(votes) > expectedVotes {
		// received enough votes, will transform to follower
		log.Println(
			"received enough votes, become a follower of:",
			appendEntryReq.LeaderId,
			", term", appendEntryReq.Term)
		m.Role = FOLLOWER
		m.Leader = appendEntryReq.LeaderId
		m.ResetTerm(appendEntryReq.Term)
		m.Server.SaveGroupNTXN(m.Group)
		return true
	} else {
		return false
	}
}
