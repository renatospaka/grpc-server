# gRPC (protoc)
## install protoc  --> this is being done by the Dockerfile
apt install -y protobuf-compiler
export PATH="$PATH:$HOME/.local/bin"
protoc --version

## instsll go packages for protoc -->
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

## command to generate the contract for gRPC calls -->
protoc --proto_path=grpc/protofile grpc/protofile/*.proto --plugin=$(go env GOPATH)/bin/protoc-gen-go-grpc --go-grpc_out=. --go_out=.

## execute a gRPC server via Evans
evans -r repl -p=50056
call CreateCourse()