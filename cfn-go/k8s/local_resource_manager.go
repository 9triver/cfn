package k8s

import (
	"fmt"
	"github.com/9triver/cfn/cfn-go/models"
	messages "github.com/9triver/ignis/proto/controller"
	"github.com/asynkron/protoactor-go/actor"
)

//const (
//	IMAGE_NAME      = "cfn-work-platform-python:0.0.0"
//	DOCKERFILE_PATH = "cfn-work-platform/docker-python"
//)

type LocalResourceManager struct {
}

func NewLocalResourceManager() *LocalResourceManager {
	return &LocalResourceManager{}
}

//func (rm *LocalResourceManager) getImage() {
//
//}

func (rm *LocalResourceManager) CreateFuncActor(appendPyFunc *messages.AppendPyFunc) *actor.PID {
	//panic("implement me")
	//client := getK8sClient()
	//node := getK8sNode(client, "epyc")
	//memory := node.Status.Capacity.Memory()
	//fmt.Println("memory:", memory)
	//return &actor.PID{}
}

func (rm *LocalResourceManager) GetResourceStatus() *models.Resource {
	//client := getK8sClient()
	return nil
}
