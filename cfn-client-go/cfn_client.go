package client

import (
	"fmt"
	"github.com/asynkron/protoactor-go/actor"
	"reflect"
	"runtime"
)

type CFNClient struct {
	remotePID *actor.PID
	imagePath string
}

func NewCFNClient(remotePID *actor.PID, imagePath string) *CFNClient {
	return &CFNClient{
		remotePID: remotePID,
		imagePath: imagePath,
	}
}

func (client *CFNClient) Submit(functionName string, codePath string) {
	// TODO: 将函数名和函数源码上传至远程 Actor

	//	// TODO: 打包创建镜像的时候需要包含最小依赖
	//	function(nil)
}
