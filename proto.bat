cd rpc/
protoc --go_out=. --go-grpc_out=. message/*.proto