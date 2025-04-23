package main

import (
	"github.com/9triver/cfn/work-platform/docker-python/workenv"
	"os"
)

func main() {
	// 1. 获取环境变量
	host := os.Getenv("ACTOR_HOST")
	port := os.Getenv("ACTOR_PORT")
	if host == "" {
		host = "0.0.0.0"
	}
	if port == "" {
		port = "8667"
	}
	_ = workenv.NewController(host, port) // 启动 grpc 服务器

	select {}
}
