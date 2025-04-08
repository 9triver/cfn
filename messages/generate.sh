# bug 的根源：没有加 --proto_path=. \
protoc --proto_path=./proto \
       --proto_path=$GOPATH/pkg/mod/  \
       --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       ./proto/*.proto