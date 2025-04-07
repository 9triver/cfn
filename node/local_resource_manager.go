package node

import (
	"github.com/asynkron/protoactor-go/actor"
	"github.com/kekwy/cfn/models"
)

type LocalResourceManager interface {
	CreateWorkerActor(requiredResource *models.Resource) *actor.PID
}
