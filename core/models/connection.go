package models

import "github.com/asynkron/protoactor-go/actor"

type NeighborConnection struct {
	PID *actor.PID
}

func NewNeighborConnection(pid *actor.PID) *NeighborConnection {
	return &NeighborConnection{PID: pid}
}
