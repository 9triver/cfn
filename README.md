# 算力网络设计与实现 示例

> Kekwy 2025.3.17
>
> kekwy@gmail.com

## 环境配置

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
```

```shell
go get github.com/asynkron/protoactor-go/...
cd $GOPATH/src/github.com/asynkron/protoactor-go
go get ./...
make
```

## 概念设计