package main

import (
	"fmt"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/remote"
)

func main() {
	fmt.Println("TaskActor: Hello World")

	// 启动远程
	system := actor.NewActorSystem()
	remoteConfig := remote.Configure("0.0.0.0", 27272)
	remoteInstance := remote.NewRemote(system, remoteConfig)
	remoteInstance.Start()

	// 注册 TaskActor
	remoteInstance.Register("function", actor.PropsFromProducer(func() actor.Actor {
		return &TaskActor{}
	}))

	select {}
}
