package server

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/dgraph-io/badger/v3"
	"github.com/hashicorp/raft"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/amjadjibon/dkv/fsm"
	"github.com/amjadjibon/dkv/proto"
)

type DKVServerV1 struct {
	proto.UnimplementedDKVServer
	db   *badger.DB
	raft *raft.Raft
}

func (d *DKVServerV1) KVGet(ctx context.Context, request *proto.KVGetRequest) (*proto.KVGetResponse, error) {
	if ctx.Err() == context.Canceled {
		return nil, ctx.Err()
	}
	if ctx.Err() == context.DeadlineExceeded {
		return nil, ctx.Err()
	}

	var txn = d.db.NewTransaction(false)
	defer func() {
		_ = txn.Commit()
	}()

	var key = []byte(request.Key)
	item, err := txn.Get(key)
	if err != nil {
		return nil, err
	}

	var value = make([]byte, 0)
	err = item.Value(func(val []byte) error {
		value = append(value, val...)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &proto.KVGetResponse{Value: value}, nil
}

func (d *DKVServerV1) KVPut(ctx context.Context, request *proto.KVPutRequest) (*emptypb.Empty, error) {
	if ctx.Err() == context.Canceled {
		return nil, ctx.Err()
	}
	if ctx.Err() == context.DeadlineExceeded {
		return nil, ctx.Err()
	}

	request.Key = strings.TrimSpace(request.Key)
	if len(request.Key) == 0 {
		return nil, errors.New("key is empty")
	}

	if d.raft.State() != raft.Leader {
		return nil, errors.New("node is not leader")
	}

	var cmd = &fsm.Command{
		Op:    "SET",
		Key:   request.Key,
		Value: request.Value,
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return nil, err
	}

	applyFuture := d.raft.Apply(data, 500*time.Second)
	err = applyFuture.Error()
	if err != nil {
		return nil, err
	}

	_, ok := applyFuture.Response().(fsm.Response)
	if !ok {
		return nil, errors.New("response not ok")
	}

	return nil, nil
}

func (d *DKVServerV1) KVDelete(ctx context.Context, request *proto.KVDeleteRequest) (*emptypb.Empty, error) {
	if ctx.Err() == context.Canceled {
		return nil, ctx.Err()
	}
	if ctx.Err() == context.DeadlineExceeded {
		return nil, ctx.Err()
	}

	request.Key = strings.TrimSpace(request.Key)
	if len(request.Key) == 0 {
		return nil, errors.New("key is empty")
	}

	if d.raft.State() != raft.Leader {
		return nil, errors.New("node is not leader")
	}

	var cmd = &fsm.Command{
		Op:    "SET",
		Key:   request.Key,
		Value: nil,
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return nil, err
	}

	applyFuture := d.raft.Apply(data, 500*time.Second)
	err = applyFuture.Error()
	if err != nil {
		return nil, err
	}

	_, ok := applyFuture.Response().(fsm.Response)
	if !ok {
		return nil, errors.New("response not ok")
	}

	return nil, nil
}

func (d *DKVServerV1) RaftJoin(ctx context.Context, request *proto.RaftJoinRequest) (*emptypb.Empty, error) {
	if ctx.Err() == context.Canceled {
		return nil, ctx.Err()
	}
	if ctx.Err() == context.DeadlineExceeded {
		return nil, ctx.Err()
	}

	if d.raft.State() != raft.Leader {
		return nil, errors.New("node is not leader")
	}

	raftConfig := d.raft.GetConfiguration()
	err := raftConfig.Error()
	if err != nil {
		return nil, err
	}

	indexFuture := d.raft.AddVoter(raft.ServerID(request.NodeId), raft.ServerAddress(request.NodeAddress), 0, 0)

	err = indexFuture.Error()
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (d *DKVServerV1) RaftLeave(ctx context.Context, request *proto.RaftLeaveRequest) (*emptypb.Empty, error) {
	if ctx.Err() == context.Canceled {
		return nil, ctx.Err()
	}
	if ctx.Err() == context.DeadlineExceeded {
		return nil, ctx.Err()
	}

	if d.raft.State() != raft.Leader {
		return nil, errors.New("node is not leader")
	}

	raftConfig := d.raft.GetConfiguration()
	err := raftConfig.Error()
	if err != nil {
		return nil, err
	}

	indexFuture := d.raft.RemoveServer(raft.ServerID(request.NodeId), 0, 0)
	err = indexFuture.Error()
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (d *DKVServerV1) RaftState(ctx context.Context, _ *emptypb.Empty) (*proto.RaftStateResponse, error) {
	if ctx.Err() == context.Canceled {
		return nil, ctx.Err()
	}
	if ctx.Err() == context.DeadlineExceeded {
		return nil, ctx.Err()
	}
	return &proto.RaftStateResponse{RaftState: d.raft.State().String()}, nil
}

func New(db *badger.DB, raft *raft.Raft) *DKVServerV1 {
	return &DKVServerV1{
		db:   db,
		raft: raft,
	}
}
