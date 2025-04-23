package main

import (
	"fmt"
	"github.com/9triver/cfn/work-platform/docker-python/workenv"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/remote"
	"os"
	"strconv"
)

func main() {
	// 1. 获取环境变量
	host := os.Getenv("ACTOR_HOST")
	port := os.Getenv("ACTOR_PORT")
	name := os.Getenv("ACTOR_NAME")
	if host == "" {
		host = "0.0.0.0"
	}
	if port == "" {
		port = "8667"
	}
	if name == "" {
		name = "cfn-work-platform"
	}
	portNum, err1 := strconv.Atoi(port)
	if err1 != nil {
		panic(err1)
	}
	// 2. 启动远程
	system := actor.NewActorSystem()
	remoteConfig := remote.Configure(host, portNum)
	//remoteConfig.AdvertisedHost = "127.0.0.1" // TODO: fix it
	//remote.ConfigOption
	remoteInstance := remote.NewRemote(system, remoteConfig)
	remoteInstance.Start()

	// 3. 启动 Actor
	context := system.Root
	props := actor.PropsFromProducer(func() actor.Actor {
		return workenv.NewController()
	})
	pid, err2 := context.SpawnNamed(props, name)

	fmt.Println(pid)

	if err2 != nil { // 定名称冲突
		context.Spawn(props)
	}
	//pid := context.Spawn(props)
	//test(pid, context)
	select {}
}
