# grpc

## go generate

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

protoc --go_out=./ --go-grpc_out=./ proto/*.proto
