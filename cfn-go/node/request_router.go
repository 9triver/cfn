package node

import (
	"strings"
	"sync"
)

type RouterEntry struct {
	Address string
	Id      string
}

// TODO: 实现对网络时延等通信开销的感知，并将其作为执行路由策略时的考虑因素之一

type RequestRouter struct {
	node          *CFNNode
	neighbors     map[RouterEntry]*HeadNodeClient
	neighborsLock *sync.Mutex
}

func NewRequestRouter(node *CFNNode) *RequestRouter {
	res := &RequestRouter{
		node:          node,
		neighbors:     make(map[RouterEntry]*HeadNodeClient),
		neighborsLock: new(sync.Mutex),
	}
	return res
}

func (router *RequestRouter) addNeighbor(pid string) {
	if len(pid) == 0 {
		return
	} else {
		parts := strings.Split(pid, "@")
		if len(parts) != 2 {
			panic("无效的 PID 字符串格式")
		}
		id := parts[0]      // Actor ID
		address := parts[1] // Actor 地址
		router.neighborsLock.Lock()
		router.neighbors[RouterEntry{address, id}] = nil
		router.neighborsLock.Unlock()
	}
}
