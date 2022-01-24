package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/dgraph-io/badger/v3"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/hashicorp/raft"
	"google.golang.org/grpc"

	"github.com/amjadjibon/dkv/proto"
	"github.com/amjadjibon/dkv/server"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	proto.RegisterDKVServer(grpcServer, server.New(&badger.DB{}, &raft.Raft{}))

	go func() {
		err = grpcServer.Serve(lis)
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		mux := runtime.NewServeMux()
		err = proto.RegisterDKVHandlerFromEndpoint(context.TODO(), mux, ":8080", nil)
		if err != nil {
			return
		}
		srv := &http.Server{
			Addr:    fmt.Sprintf(":%d", 9090),
			Handler: mux,
		}
		err = srv.ListenAndServe()
		if err != nil {
			return
		}
	}()
}
