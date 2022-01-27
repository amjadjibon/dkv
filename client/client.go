package client

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/amjadjibon/dkv/proto"
)

type DKVClient struct {
	host    string
	port    string
	timeout time.Duration
}

func (d *DKVClient) GetClient() (proto.DKVClient, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), d.timeout)
	defer cancel()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithBlock())

	var target = fmt.Sprintf("%s:%s", d.host, d.port)
	var conn, err = grpc.DialContext(ctx, target, opts...)
	if err != nil {
		return nil, err
	}

	return proto.NewDKVClient(conn), nil
}

func (d *DKVClient) GetValue(ctx context.Context, key string) ([]byte, error) {
	var client, err = d.GetClient()
	if err != nil {
		return nil, err
	}

	res, err := client.KVGet(ctx, &proto.KVGetRequest{Key: key})
	if err != nil {
		return nil, err
	}

	return res.Value, nil
}

func (d *DKVClient) SetValue(ctx context.Context, key string, value []byte) error {
	var client, err = d.GetClient()
	if err != nil {
		return err
	}

	_, err = client.KVPut(ctx, &proto.KVPutRequest{Key: key, Value: value})
	if err != nil {
		return err
	}

	return nil
}

func (d *DKVClient) DeleteValue(ctx context.Context, key string) error {
	var client, err = d.GetClient()
	if err != nil {
		return err
	}

	_, err = client.KVDelete(ctx, &proto.KVDeleteRequest{Key: key})
	if err != nil {
		return err
	}

	return nil
}

func (d *DKVClient) RaftJoin(ctx context.Context, nodeId, nodeAddress string) error {
	var client, err = d.GetClient()
	if err != nil {
		return err
	}

	_, err = client.RaftJoin(ctx, &proto.RaftJoinRequest{NodeId: nodeId, NodeAddress: nodeAddress})
	if err != nil {
		return err
	}

	return nil
}

func (d *DKVClient) RaftLeave(ctx context.Context, nodeId string) error {
	var client, err = d.GetClient()
	if err != nil {
		return err
	}

	_, err = client.RaftLeave(ctx, &proto.RaftLeaveRequest{NodeId: nodeId})
	if err != nil {
		return err
	}

	return nil
}

func (d *DKVClient) RaftState(ctx context.Context) (state string, err error) {
	client, err := d.GetClient()
	if err != nil {
		return
	}

	res, err := client.RaftState(ctx, &emptypb.Empty{})
	if err != nil {
		return
	}

	state = res.RaftState

	return state, err
}

func New(host, port string) *DKVClient {
	return &DKVClient{
		host:    host,
		port:    port,
		timeout: 5 * time.Second,
	}
}
