package k8s

import (
	"fmt"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/kekwy/cfn/models"
)

type LocalResourceManager struct {
}

func NewLocalResourceManager() *LocalResourceManager {
	return &LocalResourceManager{}
}

func (manager LocalResourceManager) CreateWorkerActor(actorImage []byte, requestResource *models.Resource) *actor.PID {
	//panic("implement me")
	client := getK8sClient()
	node := getK8sNode(client, "epyc")
	memory := node.Status.Capacity.Memory()
	fmt.Println("memory:", memory)
	return &actor.PID{}
}

func (manager LocalResourceManager) GetResourceStatus() *models.Resource {
	//client := getK8sClient()
	return nil
}
