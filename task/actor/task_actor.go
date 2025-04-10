package main

import (
	"fmt"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/kekwy/cfn/messages"
	"github.com/kekwy/cfn/task"
	"os"
	"os/exec"
	"path/filepath"
	"plugin"
)

type TaskActor struct {
	functionFileDir string
	function        task.Function
}

func (a *TaskActor) deployFunction(function *messages.TaskFunction) {
	namespace := function.GetNamespace()
	name := function.GetName()
	srcPath := filepath.Join(a.functionFileDir, "uploads", namespace, name+".go")
	soPath := filepath.Join(a.functionFileDir, "plugins", namespace, name+".so")
	// 创建必要目录
	err := os.MkdirAll("uploads", os.ModePerm)
	if err != nil {
		return
	}
	err = os.MkdirAll("plugins", os.ModePerm)
	if err != nil {
		return
	}

	// 写入源码
	err = os.WriteFile(srcPath, function.GetCode(), 0644)
	if err != nil {
		return
	}

	// 编译插件
	cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", soPath, srcPath)
	_, err = cmd.CombinedOutput()
	if err != nil {
		return
	}

	// 加载插件
	p, err := plugin.Open(soPath)
	if err != nil {
		fmt.Println("plugin open: %w", err)
	}

	sym, err := p.Lookup("GetFunction")
	if err != nil {
		fmt.Println("symbol lookup: %w", err)
	}

	factory, ok := sym.(task.FunctionFactory)
	if !ok {
		fmt.Println("invalid handler type")
	}
	a.function = factory()
}

func (a *TaskActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *actor.Started:
		fmt.Println("Started" + fmt.Sprint(msg))
	case messages.TaskFunction:
		a.deployFunction(&msg)
	}

}
