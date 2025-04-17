package main

import (
	"github.com/9triver/ignis/proto/controller"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/remote"
)

func main() {
	system := actor.NewActorSystem()
	remoteConfig := remote.Configure("localhost", 0)
	remoteInstance := remote.NewRemote(system, remoteConfig)
	remoteInstance.Start()

	root := system.Root

	root.Send(
		actor.NewPID("[::]:8667", "cfn-work-platform"),
		&controller.AppendPyFunc{
			Name:          "",
			Params:        nil,
			Venv:          "",
			Requirements:  []string{"requests==2.25.0", "numpy==1.26.4"},
			PickledObject: nil,
			Language:      0,
			Resource:      nil,
		},
	)

	select {}
}
