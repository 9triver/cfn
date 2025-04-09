package client

import (
	"fmt"
)

type Storage interface {
	Save(data string) error
}

func main() {
	fmt.Println("TaskActor: Hello World")

	//// 启动远程
	//system := actor.NewActorSystem()
	//remoteConfig := remote.Configure("0.0.0.0", 27272)
	//remoteInstance := remote.NewRemote(system, remoteConfig)
	//remoteInstance.Start()
	//
	//// 注册 TaskActor
	//remoteInstance.Register("function", actor.PropsFromProducer(func() actor.Actor {
	//	return &TaskActor{}
	//}))

	//client := NewCFNClient(nil, "./image")
	//client.Submit(func(param map[string]interface{}) map[string]interface{} {
	//	fmt.Println(param)
	//	return nil
	//})

	select {}
}
