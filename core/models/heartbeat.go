package models

import "github.com/asynkron/protoactor-go/actor"

type Heartbeat struct {
	Pid      *actor.PID
	Resource Resource
}
