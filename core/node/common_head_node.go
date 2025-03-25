package node

import (
	"fmt"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/remote"
	"github.com/kekwy/cfn/core/models"
	"strings"
)

type PIDEntry struct {
	Address string
	Id      string
}

func (p PIDEntry) toPID() *actor.PID {
	return actor.NewPID(p.Address, p.Id)
}

type CommonHeadNode struct {
	neighbors map[PIDEntry]*HeadNodeClient
	pid       *actor.PID
}

func (node *CommonHeadNode) init(context actor.Context) {
	// 获取当前结点的 pid
	node.pid = context.Self()
	// 连接邻居 head 结点
	for neighborsPID := range node.neighbors {
		context.Send(neighborsPID.toPID(),
			models.NewNeighborConnection(
				node.pid,
			))
	}
	// 创建心跳检测线程
}

func (node *CommonHeadNode) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *actor.Started:
		node.init(context) // 初始化
	case models.Heartbeat:
		fmt.Printf("Hello %v\n", msg.Pid)
	}
}

func NewCommonHeadNode(neighborPIDs []string) *CommonHeadNode {
	res := &CommonHeadNode{}
	res.neighbors = make(map[PIDEntry]*HeadNodeClient)
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

func RunCommonHeadNode(host string, port int, name string, neighborPIDs []string) *actor.PID {
	// 启动远程
	system := actor.NewActorSystem()
	remoteConfig := remote.Configure(host, port)
	remoteInstance := remote.NewRemote(system, remoteConfig)
	remoteInstance.Start()

	// 创建 Actor
	context := system.Root
	props := actor.PropsFromProducer(func() actor.Actor {
		return NewCommonHeadNode(neighborPIDs)
	})
	pid, err := context.SpawnNamed(props, name)
	if err != nil { // TODO: 指定名称冲突
		pid = context.Spawn(props)
	}
	return pid
}

type CommonHeadNodeClient struct {
}
