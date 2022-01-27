package server

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/dgraph-io/badger/v3"
	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb"
	"google.golang.org/grpc"

	"github.com/amjadjibon/dkv/conf"
	"github.com/amjadjibon/dkv/fsm"
	"github.com/amjadjibon/dkv/proto"
)

const (
	// The maxPool controls how many connections we will pool.
	maxPool = 3

	// The timeout is used to apply I/O deadlines. For InstallSnapshot, we multiply
	// the timeout by (SnapshotSize / TimeoutScale).
	// https://github.com/hashicorp/raft/blob/v1.1.2/net_transport.go#L177-L181
	tcpTimeout = 10 * time.Second

	// The `retain` parameter controls how many
	// snapshots are retained. Must be at least 1.
	raftSnapShotRetain = 2

	// raftLogCacheSize is the maximum number of logs to cache in-memory.
	// This is used to reduce disk I/O for the recently committed entries.
	raftLogCacheSize = 512
)

func BadgerRaftSetup() {
	badgerOpt := badger.DefaultOptions(conf.EnvironmentConfigIns().RaftVolumeDir)
	badgerDB, err := badger.Open(badgerOpt)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = badgerDB.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	var raftBinAddr = fmt.Sprintf("localhost:%d", conf.EnvironmentConfigIns().RaftPort)
	var raftConf = raft.DefaultConfig()
	raftConf.LocalID = raft.ServerID(conf.EnvironmentConfigIns().RaftNodeId)
	raftConf.SnapshotThreshold = 1024

	var fsmStore = fsm.New(badgerDB)

	store, err := raftboltdb.NewBoltStore(filepath.Join(conf.EnvironmentConfigIns().RaftVolumeDir, "raft.dataRepo"))
	if err != nil {
		panic(err)
	}

	// Wrap the store in a LogCache to improve performance.
	cacheStore, err := raft.NewLogCache(raftLogCacheSize, store)
	if err != nil {
		panic(err)
	}

	snapshotStore, err := raft.NewFileSnapshotStore(conf.EnvironmentConfigIns().RaftVolumeDir, raftSnapShotRetain, os.Stdout)
	if err != nil {
		panic(err)
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp", raftBinAddr)
	if err != nil {
		panic(err)
	}

	transport, err := raft.NewTCPTransport(raftBinAddr, tcpAddr, maxPool, tcpTimeout, os.Stdout)
	if err != nil {
		panic(err)
	}

	raftServer, err := raft.NewRaft(raftConf, fsmStore, cacheStore, store, snapshotStore, transport)
	if err != nil {
		panic(err)
	}

	// always start single server as a leader
	var configuration = raft.Configuration{
		Servers: []raft.Server{
			{
				ID:      raft.ServerID(conf.EnvironmentConfigIns().RaftNodeId),
				Address: transport.LocalAddr(),
			},
		},
	}

	raftServer.BootstrapCluster(configuration)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.EnvironmentConfigIns().GRPCServerPort))
	if err != nil {
		panic(err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	proto.RegisterDKVServer(grpcServer, New(badgerDB, raftServer))

	fmt.Println("rpc server started at: ", lis.Addr().String())
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
