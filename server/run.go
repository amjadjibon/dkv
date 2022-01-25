package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/dgraph-io/badger/v3"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/hashicorp/raft"
	"google.golang.org/grpc"

	"github.com/amjadjibon/dkv/proto"
)

func Run(db *badger.DB, raft *raft.Raft) {
	fmt.Println("run")
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	proto.RegisterDKVServer(grpcServer, New(db, raft))

	var mux = runtime.NewServeMux()
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}

	err = proto.RegisterDKVHandlerFromEndpoint(context.TODO(), mux, ":8080", dialOpts)
	if err != nil {
		return
	}

	// handle signal
	idleChan := make(chan struct{})

	go func() {
		signChan := make(chan os.Signal, 1)
		signal.Notify(signChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
		grpcServer.GracefulStop()
		close(idleChan)
	}()

	go func() {
		err = grpcServer.Serve(lis)
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		err = http.ListenAndServe(":9090", mux)
		if err != nil {
			panic(err)
		}
	}()

	fmt.Println("=================")

	// Blocking until the shutdown is complete
	<-idleChan
}
