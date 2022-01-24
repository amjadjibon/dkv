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
