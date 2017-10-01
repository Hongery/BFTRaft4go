// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/bftraft.proto

/*
Package bftraft is a generated protocol buffer package.

It is generated from these files:
	proto/bftraft.proto

It has these top-level messages:
	CommandRequest
	CommandResponse
	LogEntry
	RequestVoteRequest
	RequestVoteResponse
	AppendEntriesRequest
	AppendEntriesResponse
	Peer
	Node
	RaftGroup
*/
package bftraft

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CommandRequest struct {
	Group     uint64 `protobuf:"varint,1,opt,name=group" json:"group,omitempty"`
	ClientId  uint64 `protobuf:"varint,2,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	RequestId uint64 `protobuf:"varint,3,opt,name=request_id,json=requestId" json:"request_id,omitempty"`
	FuncId    uint64 `protobuf:"varint,4,opt,name=func_id,json=funcId" json:"func_id,omitempty"`
	Signature []byte `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	Arg       []byte `protobuf:"bytes,6,opt,name=arg,proto3" json:"arg,omitempty"`
}

func (m *CommandRequest) Reset()                    { *m = CommandRequest{} }
func (m *CommandRequest) String() string            { return proto.CompactTextString(m) }
func (*CommandRequest) ProtoMessage()               {}
func (*CommandRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CommandRequest) GetGroup() uint64 {
	if m != nil {
		return m.Group
	}
	return 0
}

func (m *CommandRequest) GetClientId() uint64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *CommandRequest) GetRequestId() uint64 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *CommandRequest) GetFuncId() uint64 {
	if m != nil {
		return m.FuncId
	}
	return 0
}

func (m *CommandRequest) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *CommandRequest) GetArg() []byte {
	if m != nil {
		return m.Arg
	}
	return nil
}

type CommandResponse struct {
	Group     uint64 `protobuf:"varint,1,opt,name=group" json:"group,omitempty"`
	LeaderId  uint64 `protobuf:"varint,2,opt,name=leader_id,json=leaderId" json:"leader_id,omitempty"`
	NodeId    uint64 `protobuf:"varint,3,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	RequestId uint64 `protobuf:"varint,4,opt,name=request_id,json=requestId" json:"request_id,omitempty"`
	Signature []byte `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	Result    []byte `protobuf:"bytes,6,opt,name=result,proto3" json:"result,omitempty"`
}

func (m *CommandResponse) Reset()                    { *m = CommandResponse{} }
func (m *CommandResponse) String() string            { return proto.CompactTextString(m) }
func (*CommandResponse) ProtoMessage()               {}
func (*CommandResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CommandResponse) GetGroup() uint64 {
	if m != nil {
		return m.Group
	}
	return 0
}

func (m *CommandResponse) GetLeaderId() uint64 {
	if m != nil {
		return m.LeaderId
	}
	return 0
}

func (m *CommandResponse) GetNodeId() uint64 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *CommandResponse) GetRequestId() uint64 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *CommandResponse) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *CommandResponse) GetResult() []byte {
	if m != nil {
		return m.Result
	}
	return nil
}

type LogEntry struct {
	Term    uint64          `protobuf:"varint,1,opt,name=term" json:"term,omitempty"`
	Hash    []byte          `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Command *CommandRequest `protobuf:"bytes,3,opt,name=command" json:"command,omitempty"`
}

func (m *LogEntry) Reset()                    { *m = LogEntry{} }
func (m *LogEntry) String() string            { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()               {}
func (*LogEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LogEntry) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *LogEntry) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *LogEntry) GetCommand() *CommandRequest {
	if m != nil {
		return m.Command
	}
	return nil
}

type RequestVoteRequest struct {
	Group       uint64 `protobuf:"varint,1,opt,name=group" json:"group,omitempty"`
	Term        uint64 `protobuf:"varint,2,opt,name=term" json:"term,omitempty"`
	LogIndex    uint64 `protobuf:"varint,3,opt,name=log_index,json=logIndex" json:"log_index,omitempty"`
	LogTerm     uint64 `protobuf:"varint,4,opt,name=log_term,json=logTerm" json:"log_term,omitempty"`
	CandidateId uint64 `protobuf:"varint,5,opt,name=candidate_id,json=candidateId" json:"candidate_id,omitempty"`
	Signature   []byte `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *RequestVoteRequest) Reset()                    { *m = RequestVoteRequest{} }
func (m *RequestVoteRequest) String() string            { return proto.CompactTextString(m) }
func (*RequestVoteRequest) ProtoMessage()               {}
func (*RequestVoteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RequestVoteRequest) GetGroup() uint64 {
	if m != nil {
		return m.Group
	}
	return 0
}

func (m *RequestVoteRequest) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *RequestVoteRequest) GetLogIndex() uint64 {
	if m != nil {
		return m.LogIndex
	}
	return 0
}

func (m *RequestVoteRequest) GetLogTerm() uint64 {
	if m != nil {
		return m.LogTerm
	}
	return 0
}

func (m *RequestVoteRequest) GetCandidateId() uint64 {
	if m != nil {
		return m.CandidateId
	}
	return 0
}

func (m *RequestVoteRequest) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type RequestVoteResponse struct {
	Group       uint64 `protobuf:"varint,1,opt,name=group" json:"group,omitempty"`
	Term        uint64 `protobuf:"varint,2,opt,name=term" json:"term,omitempty"`
	NodeId      uint64 `protobuf:"varint,3,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	CandidateId uint64 `protobuf:"varint,4,opt,name=candidate_id,json=candidateId" json:"candidate_id,omitempty"`
	Granted     bool   `protobuf:"varint,5,opt,name=granted" json:"granted,omitempty"`
	Signature   []byte `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *RequestVoteResponse) Reset()                    { *m = RequestVoteResponse{} }
func (m *RequestVoteResponse) String() string            { return proto.CompactTextString(m) }
func (*RequestVoteResponse) ProtoMessage()               {}
func (*RequestVoteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RequestVoteResponse) GetGroup() uint64 {
	if m != nil {
		return m.Group
	}
	return 0
}

func (m *RequestVoteResponse) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *RequestVoteResponse) GetNodeId() uint64 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *RequestVoteResponse) GetCandidateId() uint64 {
	if m != nil {
		return m.CandidateId
	}
	return 0
}

func (m *RequestVoteResponse) GetGranted() bool {
	if m != nil {
		return m.Granted
	}
	return false
}

func (m *RequestVoteResponse) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type AppendEntriesRequest struct {
	Group        uint64                 `protobuf:"varint,1,opt,name=group" json:"group,omitempty"`
	Term         uint64                 `protobuf:"varint,2,opt,name=term" json:"term,omitempty"`
	LeaderId     uint64                 `protobuf:"varint,3,opt,name=leader_id,json=leaderId" json:"leader_id,omitempty"`
	PrevLogIndex uint64                 `protobuf:"varint,4,opt,name=prev_log_index,json=prevLogIndex" json:"prev_log_index,omitempty"`
	PrevLogTerm  uint64                 `protobuf:"varint,5,opt,name=prev_log_term,json=prevLogTerm" json:"prev_log_term,omitempty"`
	Signature    []byte                 `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
	QuorumVotes  []*RequestVoteResponse `protobuf:"bytes,7,rep,name=quorum_votes,json=quorumVotes" json:"quorum_votes,omitempty"`
	Entries      []*LogEntry            `protobuf:"bytes,8,rep,name=entries" json:"entries,omitempty"`
}

func (m *AppendEntriesRequest) Reset()                    { *m = AppendEntriesRequest{} }
func (m *AppendEntriesRequest) String() string            { return proto.CompactTextString(m) }
func (*AppendEntriesRequest) ProtoMessage()               {}
func (*AppendEntriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *AppendEntriesRequest) GetGroup() uint64 {
	if m != nil {
		return m.Group
	}
	return 0
}

func (m *AppendEntriesRequest) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *AppendEntriesRequest) GetLeaderId() uint64 {
	if m != nil {
		return m.LeaderId
	}
	return 0
}

func (m *AppendEntriesRequest) GetPrevLogIndex() uint64 {
	if m != nil {
		return m.PrevLogIndex
	}
	return 0
}

func (m *AppendEntriesRequest) GetPrevLogTerm() uint64 {
	if m != nil {
		return m.PrevLogTerm
	}
	return 0
}

func (m *AppendEntriesRequest) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *AppendEntriesRequest) GetQuorumVotes() []*RequestVoteResponse {
	if m != nil {
		return m.QuorumVotes
	}
	return nil
}

func (m *AppendEntriesRequest) GetEntries() []*LogEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type AppendEntriesResponse struct {
	Group     uint64 `protobuf:"varint,1,opt,name=group" json:"group,omitempty"`
	Term      uint64 `protobuf:"varint,2,opt,name=term" json:"term,omitempty"`
	Index     uint64 `protobuf:"varint,3,opt,name=index" json:"index,omitempty"`
	NodeId    uint64 `protobuf:"varint,4,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	Successed bool   `protobuf:"varint,5,opt,name=successed" json:"successed,omitempty"`
	Convinced bool   `protobuf:"varint,6,opt,name=convinced" json:"convinced,omitempty"`
	Hash      []byte `protobuf:"bytes,7,opt,name=hash,proto3" json:"hash,omitempty"`
	Signature []byte `protobuf:"bytes,8,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *AppendEntriesResponse) Reset()                    { *m = AppendEntriesResponse{} }
func (m *AppendEntriesResponse) String() string            { return proto.CompactTextString(m) }
func (*AppendEntriesResponse) ProtoMessage()               {}
func (*AppendEntriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AppendEntriesResponse) GetGroup() uint64 {
	if m != nil {
		return m.Group
	}
	return 0
}

func (m *AppendEntriesResponse) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *AppendEntriesResponse) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *AppendEntriesResponse) GetNodeId() uint64 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *AppendEntriesResponse) GetSuccessed() bool {
	if m != nil {
		return m.Successed
	}
	return false
}

func (m *AppendEntriesResponse) GetConvinced() bool {
	if m != nil {
		return m.Convinced
	}
	return false
}

func (m *AppendEntriesResponse) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *AppendEntriesResponse) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Peer struct {
	Id   uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Host uint64 `protobuf:"varint,2,opt,name=host" json:"host,omitempty"`
}

func (m *Peer) Reset()                    { *m = Peer{} }
func (m *Peer) String() string            { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()               {}
func (*Peer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Peer) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Peer) GetHost() uint64 {
	if m != nil {
		return m.Host
	}
	return 0
}

type Node struct {
	Id        uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	LastSeen  uint64 `protobuf:"varint,2,opt,name=last_seen,json=lastSeen" json:"last_seen,omitempty"`
	Weight    uint32 `protobuf:"varint,3,opt,name=weight" json:"weight,omitempty"`
	Online    bool   `protobuf:"varint,4,opt,name=online" json:"online,omitempty"`
	PublicKey string `protobuf:"bytes,5,opt,name=public_key,json=publicKey" json:"public_key,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Node) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Node) GetLastSeen() uint64 {
	if m != nil {
		return m.LastSeen
	}
	return 0
}

func (m *Node) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Node) GetOnline() bool {
	if m != nil {
		return m.Online
	}
	return false
}

func (m *Node) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

type RaftGroup struct {
	Replications uint32 `protobuf:"varint,1,opt,name=replications" json:"replications,omitempty"`
	Id           uint64 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	LeaderPeer   uint64 `protobuf:"varint,3,opt,name=leader_peer,json=leaderPeer" json:"leader_peer,omitempty"`
}

func (m *RaftGroup) Reset()                    { *m = RaftGroup{} }
func (m *RaftGroup) String() string            { return proto.CompactTextString(m) }
func (*RaftGroup) ProtoMessage()               {}
func (*RaftGroup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *RaftGroup) GetReplications() uint32 {
	if m != nil {
		return m.Replications
	}
	return 0
}

func (m *RaftGroup) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RaftGroup) GetLeaderPeer() uint64 {
	if m != nil {
		return m.LeaderPeer
	}
	return 0
}

func init() {
	proto.RegisterType((*CommandRequest)(nil), "bftraft.CommandRequest")
	proto.RegisterType((*CommandResponse)(nil), "bftraft.CommandResponse")
	proto.RegisterType((*LogEntry)(nil), "bftraft.LogEntry")
	proto.RegisterType((*RequestVoteRequest)(nil), "bftraft.RequestVoteRequest")
	proto.RegisterType((*RequestVoteResponse)(nil), "bftraft.RequestVoteResponse")
	proto.RegisterType((*AppendEntriesRequest)(nil), "bftraft.AppendEntriesRequest")
	proto.RegisterType((*AppendEntriesResponse)(nil), "bftraft.AppendEntriesResponse")
	proto.RegisterType((*Peer)(nil), "bftraft.Peer")
	proto.RegisterType((*Node)(nil), "bftraft.Node")
	proto.RegisterType((*RaftGroup)(nil), "bftraft.RaftGroup")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BFTRaft service

type BFTRaftClient interface {
	ExecCommand(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandResponse, error)
	RequestVote(ctx context.Context, in *RequestVoteRequest, opts ...grpc.CallOption) (*RequestVoteResponse, error)
	AppendEntries(ctx context.Context, in *AppendEntriesRequest, opts ...grpc.CallOption) (*AppendEntriesResponse, error)
}

type bFTRaftClient struct {
	cc *grpc.ClientConn
}

func NewBFTRaftClient(cc *grpc.ClientConn) BFTRaftClient {
	return &bFTRaftClient{cc}
}

func (c *bFTRaftClient) ExecCommand(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandResponse, error) {
	out := new(CommandResponse)
	err := grpc.Invoke(ctx, "/bftraft.BFTRaft/ExecCommand", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFTRaftClient) RequestVote(ctx context.Context, in *RequestVoteRequest, opts ...grpc.CallOption) (*RequestVoteResponse, error) {
	out := new(RequestVoteResponse)
	err := grpc.Invoke(ctx, "/bftraft.BFTRaft/RequestVote", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFTRaftClient) AppendEntries(ctx context.Context, in *AppendEntriesRequest, opts ...grpc.CallOption) (*AppendEntriesResponse, error) {
	out := new(AppendEntriesResponse)
	err := grpc.Invoke(ctx, "/bftraft.BFTRaft/AppendEntries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BFTRaft service

type BFTRaftServer interface {
	ExecCommand(context.Context, *CommandRequest) (*CommandResponse, error)
	RequestVote(context.Context, *RequestVoteRequest) (*RequestVoteResponse, error)
	AppendEntries(context.Context, *AppendEntriesRequest) (*AppendEntriesResponse, error)
}

func RegisterBFTRaftServer(s *grpc.Server, srv BFTRaftServer) {
	s.RegisterService(&_BFTRaft_serviceDesc, srv)
}

func _BFTRaft_ExecCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFTRaftServer).ExecCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bftraft.BFTRaft/ExecCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFTRaftServer).ExecCommand(ctx, req.(*CommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFTRaft_RequestVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFTRaftServer).RequestVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bftraft.BFTRaft/RequestVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFTRaftServer).RequestVote(ctx, req.(*RequestVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFTRaft_AppendEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFTRaftServer).AppendEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bftraft.BFTRaft/AppendEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFTRaftServer).AppendEntries(ctx, req.(*AppendEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BFTRaft_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bftraft.BFTRaft",
	HandlerType: (*BFTRaftServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecCommand",
			Handler:    _BFTRaft_ExecCommand_Handler,
		},
		{
			MethodName: "RequestVote",
			Handler:    _BFTRaft_RequestVote_Handler,
		},
		{
			MethodName: "AppendEntries",
			Handler:    _BFTRaft_AppendEntries_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/bftraft.proto",
}

func init() { proto.RegisterFile("proto/bftraft.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 744 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xae, 0x13, 0x27, 0x76, 0x4e, 0x92, 0x02, 0xd3, 0x42, 0x4d, 0x7f, 0xa0, 0x58, 0x5c, 0x54,
	0x20, 0x15, 0x51, 0x1e, 0x00, 0x51, 0x54, 0x50, 0xa0, 0x42, 0x95, 0xa9, 0xb8, 0x0d, 0xae, 0x7d,
	0xe2, 0x5a, 0x38, 0x33, 0xee, 0xcc, 0xb8, 0xb4, 0xb7, 0xbc, 0x0e, 0x77, 0xbb, 0x17, 0xfb, 0x36,
	0xfb, 0x02, 0xfb, 0x12, 0xab, 0xf9, 0xb1, 0xf3, 0x9f, 0xd5, 0xee, 0xdd, 0x9c, 0xef, 0xcc, 0x4c,
	0xbe, 0xf3, 0x7d, 0x9f, 0x27, 0xb0, 0x57, 0x72, 0x26, 0xd9, 0x77, 0x77, 0x13, 0xc9, 0xe3, 0x89,
	0x3c, 0xd7, 0x15, 0xf1, 0x6c, 0x19, 0xfe, 0xef, 0xc0, 0xee, 0xcf, 0x6c, 0x3a, 0x8d, 0x69, 0x1a,
	0xe1, 0x43, 0x85, 0x42, 0x92, 0x7d, 0xe8, 0x64, 0x9c, 0x55, 0x65, 0xe0, 0x9c, 0x3a, 0x67, 0x6e,
	0x64, 0x0a, 0x72, 0x04, 0xbd, 0xa4, 0xc8, 0x91, 0xca, 0x71, 0x9e, 0x06, 0x2d, 0xdd, 0xf1, 0x0d,
	0x30, 0x4a, 0xc9, 0x09, 0x00, 0x37, 0xa7, 0x55, 0xb7, 0xad, 0xbb, 0x3d, 0x8b, 0x8c, 0x52, 0x72,
	0x00, 0xde, 0xa4, 0xa2, 0x89, 0xea, 0xb9, 0xba, 0xd7, 0x55, 0xe5, 0x28, 0x25, 0xc7, 0xd0, 0x13,
	0x79, 0x46, 0x63, 0x59, 0x71, 0x0c, 0x3a, 0xa7, 0xce, 0xd9, 0x20, 0x9a, 0x01, 0xe4, 0x63, 0x68,
	0xc7, 0x3c, 0x0b, 0xba, 0x1a, 0x57, 0xcb, 0xf0, 0x85, 0x03, 0x1f, 0x35, 0x6c, 0x45, 0xc9, 0xa8,
	0xc0, 0xcd, 0x74, 0x0b, 0x8c, 0x53, 0xe4, 0x73, 0x74, 0x0d, 0x60, 0xf8, 0x50, 0x96, 0xe2, 0x8c,
	0x6b, 0x57, 0x95, 0x2b, 0x73, 0xb8, 0xcb, 0x73, 0x6c, 0xa7, 0xfb, 0x19, 0x74, 0x39, 0x8a, 0xaa,
	0x90, 0x96, 0xb1, 0xad, 0x42, 0x04, 0xff, 0x9a, 0x65, 0x57, 0x54, 0xf2, 0x67, 0x42, 0xc0, 0x95,
	0xc8, 0xa7, 0x96, 0xab, 0x5e, 0x2b, 0xec, 0x3e, 0x16, 0xf7, 0x9a, 0xe5, 0x20, 0xd2, 0x6b, 0xf2,
	0x3d, 0x78, 0x89, 0x99, 0x53, 0x33, 0xec, 0x5f, 0x1c, 0x9c, 0xd7, 0x06, 0x2e, 0xba, 0x15, 0xd5,
	0xfb, 0xc2, 0x57, 0x0e, 0x10, 0x0b, 0xfe, 0xc5, 0x24, 0x6e, 0x77, 0xb3, 0xe6, 0xd1, 0x9a, 0xe3,
	0xa1, 0x24, 0x63, 0xd9, 0x38, 0xa7, 0x29, 0x3e, 0x59, 0x5d, 0xfc, 0x82, 0x65, 0x23, 0x55, 0x93,
	0xcf, 0x41, 0xad, 0xc7, 0xfa, 0x90, 0xd1, 0xc5, 0x2b, 0x58, 0x76, 0xab, 0xce, 0x7d, 0x05, 0x83,
	0x24, 0xa6, 0x69, 0x9e, 0xc6, 0x52, 0x4b, 0xda, 0xd1, 0xed, 0x7e, 0x83, 0x2d, 0x0b, 0xd7, 0x5d,
	0x12, 0x2e, 0x7c, 0xe9, 0xc0, 0xde, 0x02, 0xf3, 0xad, 0xce, 0xae, 0xa3, 0xbe, 0xd1, 0xd0, 0x65,
	0x6e, 0xee, 0x2a, 0xb7, 0x00, 0xbc, 0x8c, 0xc7, 0x54, 0xa2, 0x61, 0xee, 0x47, 0x75, 0xf9, 0x2e,
	0xd6, 0x2d, 0xd8, 0xff, 0xa9, 0x2c, 0x91, 0xa6, 0xca, 0xda, 0x1c, 0xc5, 0x87, 0x29, 0xde, 0x84,
	0xb4, 0xbd, 0x14, 0xd2, 0xaf, 0x61, 0xb7, 0xe4, 0xf8, 0x38, 0x9e, 0x79, 0x62, 0xc8, 0x0f, 0x14,
	0x7a, 0x5d, 0xfb, 0x12, 0xc2, 0xb0, 0xd9, 0xa5, 0xef, 0xb7, 0xea, 0xdb, 0x4d, 0xda, 0xa0, 0xad,
	0x73, 0x90, 0x1f, 0x61, 0xf0, 0x50, 0x31, 0x5e, 0x4d, 0xc7, 0x8f, 0x4c, 0xa2, 0x08, 0xbc, 0xd3,
	0xf6, 0x59, 0xff, 0xe2, 0xb8, 0xc9, 0xdb, 0x1a, 0x67, 0xa2, 0xbe, 0x39, 0xa1, 0x30, 0x41, 0xbe,
	0x05, 0x0f, 0x8d, 0x02, 0x81, 0xaf, 0xcf, 0x7e, 0xd2, 0x9c, 0xad, 0x73, 0x1f, 0xd5, 0x3b, 0xc2,
	0xd7, 0x0e, 0x7c, 0xba, 0xa4, 0xda, 0x7b, 0xbb, 0xbd, 0x0f, 0x9d, 0xf9, 0x90, 0x9a, 0x62, 0x3e,
	0x03, 0xee, 0x42, 0x06, 0xd4, 0xf8, 0x55, 0x92, 0xa0, 0x10, 0x8d, 0xc5, 0x33, 0x40, 0x75, 0x13,
	0x46, 0x1f, 0x73, 0x9a, 0x60, 0xaa, 0xc5, 0xf1, 0xa3, 0x19, 0xd0, 0x7c, 0x9b, 0xde, 0xdc, 0xb7,
	0xb9, 0x20, 0xa7, 0xbf, 0x1c, 0x8b, 0x6f, 0xc0, 0xbd, 0x41, 0xe4, 0x64, 0x17, 0x5a, 0x79, 0x6a,
	0x67, 0x69, 0xe5, 0xe6, 0x26, 0x26, 0x64, 0x3d, 0x88, 0x5a, 0x87, 0xff, 0x39, 0xe0, 0xfe, 0xc1,
	0x52, 0x5c, 0xd9, 0xac, 0x82, 0x11, 0x0b, 0x39, 0x16, 0x88, 0xb4, 0x79, 0xbd, 0x62, 0x21, 0xff,
	0x44, 0xa4, 0xea, 0x9d, 0xf9, 0x17, 0xf3, 0xec, 0x5e, 0xea, 0xf9, 0x87, 0x91, 0xad, 0x14, 0xce,
	0x68, 0x91, 0x53, 0xd4, 0xf3, 0xfb, 0x91, 0xad, 0xd4, 0xa3, 0x56, 0x56, 0x77, 0x45, 0x9e, 0x8c,
	0xff, 0xc1, 0x67, 0x2d, 0x40, 0x2f, 0xea, 0x19, 0xe4, 0x77, 0x7c, 0x0e, 0xff, 0x86, 0x5e, 0x14,
	0x4f, 0xe4, 0xaf, 0x5a, 0xee, 0x10, 0x06, 0x1c, 0xcb, 0x22, 0x4f, 0x62, 0x99, 0x33, 0x2a, 0x34,
	0xa5, 0x61, 0xb4, 0x80, 0x59, 0xb2, 0xad, 0x86, 0xec, 0x97, 0xd0, 0xb7, 0x29, 0x2e, 0x11, 0xb9,
	0x35, 0x05, 0x0c, 0xa4, 0xa4, 0xb8, 0x78, 0xe3, 0x80, 0x77, 0xf9, 0xcb, 0xad, 0xfa, 0x15, 0x72,
	0x09, 0xfd, 0xab, 0x27, 0x4c, 0xec, 0x23, 0x46, 0x36, 0x3d, 0x6b, 0x87, 0xc1, 0x6a, 0xc3, 0xe4,
	0x24, 0xdc, 0x21, 0xbf, 0x41, 0x7f, 0x2e, 0x94, 0xe4, 0x68, 0x7d, 0x54, 0xcd, 0x3d, 0x5b, 0x73,
	0x1c, 0xee, 0x90, 0x1b, 0x18, 0x2e, 0xc4, 0x91, 0x9c, 0x34, 0x07, 0xd6, 0x7d, 0xdc, 0x87, 0x5f,
	0x6c, 0x6a, 0xd7, 0x37, 0xde, 0x75, 0xf5, 0x3f, 0xec, 0x0f, 0x6f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x0f, 0xd1, 0x9e, 0xf3, 0x78, 0x07, 0x00, 0x00,
}
