package node

import "github.com/asynkron/protoactor-go/actor"

type HeadNodeClient interface {
	isConnected() bool
	addNeighbor(actor.PID)
}
