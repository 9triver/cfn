package main

import "github.com/asynkron/protoactor-go/actor"

type TaskActor struct {
	isReady bool
}

func (a TaskActor) Receive(c actor.Context) {
	switch msg := c.Message().(type) {
	case *actor.Started:
		a.isReady = true
	}
}
