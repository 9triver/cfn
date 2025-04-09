package client

import (
	"fmt"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/remote"
	"github.com/kekwy/cfn/client/messages"
	"google.golang.org/protobuf/types/known/structpb"
	"os"
	"strconv"
	"sync"
	"time"
)

type Environment interface {
	DeployFunction(function TaskFunction)
	Execute()
}

type TaskActor struct {
	isReady      bool
	taskFunction TaskFunction
	mutex        sync.Mutex
}

func NewTaskActor() *TaskActor {
	return &TaskActor{
		isReady:      false,
		taskFunction: nil,
	}
}

func (a *TaskActor) Receive(c actor.Context) {
	switch msg := c.Message().(type) {
	case *actor.Started:
		a.isReady = true
	case *TaskInputs:
		if !a.isReady || a.taskFunction == nil {
			c.Respond(&messages.TaskError{
				Message: "未初始化",
			})
			return
		}
		taskResults, err := a.taskFunction.Apply(msg)
		if err != nil {
			c.Respond(&messages.TaskError{
				Message: err.Error(),
			})
			return
		}
		c.Respond(taskResults)
	}
}

type ActorEnvironment struct {
	taskActor *TaskActor
}

func (a ActorEnvironment) DeployFunction(function TaskFunction) {
	a.taskActor.taskFunction = function
}

func (a ActorEnvironment) Execute() {
	// 1. 获取环境变量
	host := os.Getenv("ACTOR_HOST")
	port := os.Getenv("ACTOR_PORT")
	name := os.Getenv("ACTOR_NAME")
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "8667"
	}
	portNum, err1 := strconv.Atoi(port)
	if err1 != nil {
		panic(err1)
	}
	// 2. 启动远程
	system := actor.NewActorSystem()
	remoteConfig := remote.Configure(host, portNum)
	remoteInstance := remote.NewRemote(system, remoteConfig)
	remoteInstance.Start()

	// 3. 启动 Actor
	context := system.Root
	props := actor.PropsFromProducer(func() actor.Actor {
		return a.taskActor
	})
	if name != "" {
		_, err2 := context.SpawnNamed(props, name)
		if err2 != nil { // 定名称冲突
			context.Spawn(props)
		}
	} else {
		context.Spawn(props)
		//pid := context.Spawn(props)
		//test(pid, context)
	}
	select {}
}

func test(pid *actor.PID, context *actor.RootContext) {
	time.Sleep(100 * time.Millisecond)

	newStruct, _ := structpb.NewStruct(map[string]interface{}{
		"nums": []any{1, 2, 3},
	})
	future := context.RequestFuture(pid, (*TaskInputs)(newStruct), 5*time.Second)
	res, err := future.Result()
	outputs := res.(*TaskOutputs)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(outputs.Fields["sum"].GetNumberValue())
	}
	future = context.RequestFuture(pid, (*TaskInputs)(newStruct), 5*time.Second)
	res, err = future.Result()
	outputs = res.(*TaskOutputs)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(outputs.Fields["sum"].GetNumberValue())
	}
}

func NewActorEnvironment() *ActorEnvironment {
	return &ActorEnvironment{
		taskActor: NewTaskActor(),
	}
}

// GetEnvironment 根据系统环境变量初始化环境
func GetEnvironment() Environment {
	envType := os.Getenv("ENVIRONMENT")
	switch envType {
	case "actor":
		return NewActorEnvironment()
	}
	return GetDefaultEnvironment()
}

func GetDefaultEnvironment() Environment {
	return NewActorEnvironment()
}
