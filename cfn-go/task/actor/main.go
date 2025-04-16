package main

import (
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
		host = "localhost"
	}
	if port == "" {
		port = "8667"
	}
	portNum, err1 := strconv.Atoi(port)
	if err1 != nil {
		panic(err1)
	}
	// 2. 启动远程
	system := actor.NewActorSystem()
	remoteConfig := remote.Configure(host, portNum)
	remoteInstance := remote.NewRemote(system, remoteConfig)
	remoteInstance.Start()

	// 3. 启动 Actor
	context := system.Root
	props := actor.PropsFromProducer(func() actor.Actor {
		return &TaskActor{}
	})
	if name != "" {
		_, err2 := context.SpawnNamed(props, name)
		if err2 != nil { // 定名称冲突
			context.Spawn(props)
		}
	} else {
		context.Spawn(props)
		//pid := context.Spawn(props)
		//test(pid, context)
	}
	select {}
}
