package models

import "github.com/asynkron/protoactor-go/actor"

type PIDMessage struct {
	PID *actor.PID
}

func (p PIDMessage) Reset() {
	//TODO implement me
	panic("implement me")
}

func (p PIDMessage) String() string {
	//TODO implement me
	panic("implement me")
}

func (p PIDMessage) ProtoMessage() {
	//TODO implement me
	panic("implement me")
}

func NewPIDMessage(pid *actor.PID) *PIDMessage {
	return &PIDMessage{PID: pid}
}
