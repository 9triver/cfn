package node

import (
	"fmt"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/remote"
	"github.com/kekwy/cfn/messages"
	"github.com/kekwy/cfn/models"
	"log/slog"
	"strings"
	"time"
)

type PIDEntry struct {
	Address string
	Id      string
}

func (p PIDEntry) toPID() *actor.PID {
	return actor.NewPID(p.Address, p.Id)
}

type HeadNode struct {

	//==================================================================================================================
	// Actor 通信
	//------------------------------------------------------------------------------------------------------------------

	neighbors map[PIDEntry]*HeadNodeClient
	pid       *actor.PID
	behavior  *actor.Behavior

	//==================================================================================================================
	// Web 通信
	//------------------------------------------------------------------------------------------------------------------

	webUrl string

	//==================================================================================================================
	// 节点状态
	//------------------------------------------------------------------------------------------------------------------

	resource *models.Resource

	//==================================================================================================================
	// 守护线程
	//------------------------------------------------------------------------------------------------------------------

	ticker *time.Ticker
}

func (node *HeadNode) unexpectedMessage(msg interface{}, logger *slog.Logger) {
	logger.Warn(fmt.Sprintf("Actor: %v. Unexpected message type. %v", node.pid, msg))
}

// 初始化
//
// initBehavior
func (node *HeadNode) initBehavior(context actor.Context) {
	logger := context.Logger()
	switch msg := context.Message().(type) {
	case *actor.Started:
		node.init(context) // 初始化
	default:
		node.unexpectedMessage(msg, logger)
	}
}

// init
func (node *HeadNode) init(context actor.Context) {
	// 获取当前结点的 pid
	node.pid = context.Self()
	// 连接邻居 head 结点
	for neighborsPID := range node.neighbors {
		context.Send(neighborsPID.toPID(),
			&messages.Echo{
				Sender: node.pid,
			})
	}
	// 创建心跳检测计时器
	context.Watch(node.pid)
	// 切换 actor 状态
	node.behavior.Become(node.runningBehavior)
}

// 运行中
//
// runningBehavior
func (node *HeadNode) runningBehavior(context actor.Context) {
	logger := context.Logger()
	switch msg := context.Message().(type) {
	case *messages.Echo:
		node.echoHandler(msg, context)
	default:
		node.unexpectedMessage(msg, logger)
	}
}

func (node *HeadNode) echoHandler(msg *messages.Echo, context actor.Context) {
	tmp := "hello!" + msg.Sender.Id
	context.Send(msg.Sender, &messages.Text{
		Text: tmp,
	})
}

func (node *HeadNode) Receive(context actor.Context) {
	node.behavior.Receive(context)

}

// ---------------------------------------------------------------------------------------------------------------------
// 创建和启动
//----------------------------------------------------------------------------------------------------------------------

func NewHeadNode(neighborPIDs []string) *HeadNode {
	res := &HeadNode{}
	res.neighbors = make(map[PIDEntry]*HeadNodeClient)
	res.behavior = &actor.Behavior{}
	res.behavior.Become(res.initBehavior)
	if neighborPIDs == nil || len(neighborPIDs) == 0 {
		return res
	}
	// string 转 PIDEntry
	// e.g. "headnode@10.0.2.1:8080"
	for _, neighborPID := range neighborPIDs {
		parts := strings.Split(neighborPID, "@")
		if len(parts) != 2 {
			panic("无效的 PID 字符串格式")
		}
		id := parts[0]      // Actor ID
		address := parts[1] // Actor 地址
		res.neighbors[PIDEntry{Id: id, Address: address}] = nil
	}
	return res
}

func RunHeadNode(host string, port int, name string, neighborPIDs []string) *actor.PID {
	// 启动远程
	system := actor.NewActorSystem()
	remoteConfig := remote.Configure(host, port)
	remoteInstance := remote.NewRemote(system, remoteConfig)
	remoteInstance.Start()

	// 创建 Actor
	context := system.Root
	props := actor.PropsFromProducer(func() actor.Actor {
		return NewHeadNode(neighborPIDs)
	})
	pid, err := context.SpawnNamed(props, name)
	if err != nil { // TODO: 指定名称冲突
		pid = context.Spawn(props)
	}
	return pid
}
