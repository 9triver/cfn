package node

import (
	"github.com/9triver/cfn/cfn-go/models"
	"github.com/asynkron/protoactor-go/actor"
)

type LocalResourceManager interface {
	CreateWorkerActor(requiredResource *models.Resource) *actor.PID
}
