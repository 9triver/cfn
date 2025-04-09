package node

import (
	"fmt"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/remote"
	"github.com/kekwy/cfn/messages"
	"log/slog"
)

type ActorCore struct {
	node     *CFNNode
	pid      *actor.PID
	behavior *actor.Behavior
}

func NewActorCore(node *CFNNode) *ActorCore {
	core := &ActorCore{
		node: node,
	}
	core.behavior = &actor.Behavior{}
	core.behavior.Become(core.initBehavior)
	return core
}

func (core *ActorCore) Start(host string, port int, name string) {
	// 启动远程
	system := actor.NewActorSystem()
	remoteConfig := remote.Configure(host, port)
	remoteInstance := remote.NewRemote(system, remoteConfig)
	remoteInstance.Start()

	// 启动 Actor
	context := system.Root
	props := actor.PropsFromProducer(func() actor.Actor {
		return core
	})
	pid, err := context.SpawnNamed(props, name)
	if err != nil { // TODO: 指定名称冲突
		pid = context.Spawn(props)
	}
	core.pid = pid
}

func (core *ActorCore) GetPID() *actor.PID {
	return core.pid
}

func (core *ActorCore) Receive(context actor.Context) {
	core.behavior.Receive(context)
}

//======================================================================================================================
// Behavior
//----------------------------------------------------------------------------------------------------------------------

func (core *ActorCore) initBehavior(context actor.Context) {
	logger := context.Logger()
	switch msg := context.Message().(type) {
	case *actor.Started:
		core.initHandler(context) // 初始化
	default:
		core.unexpectedMessageHandler(msg, logger)
	}
}

func (core *ActorCore) runningBehavior(context actor.Context) {
	logger := context.Logger()
	switch msg := context.Message().(type) {
	case *messages.Echo:
		core.echoHandler(msg, context)
	default:
		core.unexpectedMessageHandler(msg, logger)
	}
}

//======================================================================================================================
// Handler
//----------------------------------------------------------------------------------------------------------------------

func (core *ActorCore) initHandler(context actor.Context) {
	// 获取当前结点的 pid
	core.pid = context.Self()
	// 连接邻居 head 结点
	//for neighborsPID := range core.neighbors {
	//	context.Send(neighborsPID.toPID(),
	//		&messages.Echo{
	//			Sender: node.pid,
	//		})
	//}
	// 创建心跳检测计时器
	context.Watch(core.pid)
	// 切换 client 状态
	core.behavior.Become(core.runningBehavior)
}

func (core *ActorCore) echoHandler(msg *messages.Echo, context actor.Context) {
	tmp := "hello!" + msg.Sender.Id
	context.Send(msg.Sender, &messages.Text{
		Text: tmp,
	})
}

func (core *ActorCore) unexpectedMessageHandler(msg interface{}, logger *slog.Logger) {
	logger.Warn(fmt.Sprintf("Actor: %v. Unexpected message type. %v", core.pid, msg))
}
