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
	// 1. 获取函数名
	fmt.Println(runtime.FuncForPC(reflect.ValueOf(function).Pointer()).Name())

	// 2. 生成 Docker File，在镜像的环境变量中指定 TaskFunc
	//reflect.
	//	// TODO: 打包创建镜像的时候需要包含最小的依赖闭包
	//	function(nil)
}
