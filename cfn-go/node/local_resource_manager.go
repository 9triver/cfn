package node

import (
	messages "github.com/9triver/ignis/proto/controller"
	"github.com/asynkron/protoactor-go/actor"
)

type LocalResourceManager interface {
	CreateFuncActor(appendPyFunc *messages.AppendPyFunc) *actor.PID
}
