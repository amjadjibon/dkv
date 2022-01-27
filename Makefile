update:
	go get github.com/dgraph-io/badger/v3
	go get github.com/hashicorp/raft
	go get github.com/spf13/viper
	go get github.com/golang/protobuf
	go get google.golang.org/grpc
	go get github.com/grpc-ecosystem/grpc-gateway
	go get github.com/golang/protobuf/protoc-gen-go

update-protoc:
	@protoc \
		-I=./proto \
		--go_opt=module=github.com/amjadjibon/dkv \
		--go_out=. \
		--go-grpc_opt=module=github.com/amjadjibon/dkv \
		--go-grpc_out=. \
		--grpc-gateway_opt=module=github.com/amjadjibon/dkv \
		--grpc-gateway_out=logtostderr=true:. \
		--swagger_out=logtostderr=true:swagger \
		./proto/dkv.proto

run-server-1:
	GRPC_SERVER_PORT=6001 RAFT_NODE_ID=raft_node_1 RAFT_PORT=7001 RAFT_VOLUME_DIR=/home/amjad/dkvnode_data_1 go run main.go run
run-server-2:
	GRPC_SERVER_PORT=6002 RAFT_NODE_ID=raft_node_2 RAFT_PORT=7002 RAFT_VOLUME_DIR=/home/amjad/dkvnode_data_2 go run main.go run
run-server-3:
	GRPC_SERVER_PORT=6003 RAFT_NODE_ID=raft_node_3 RAFT_PORT=7003 RAFT_VOLUME_DIR=/home/amjad/dkvnode_data_3 go run main.go run
run-server-4:
	GRPC_SERVER_PORT=6004 RAFT_NODE_ID=raft_node_4 RAFT_PORT=7004 RAFT_VOLUME_DIR=/home/amjad/dkvnode_data_4 go run main.go run