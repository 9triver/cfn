package node

import "time"

type DaemonThread struct {
	ticker   *time.Ticker
	node     *CFNNode
	interval time.Duration
	stop     chan struct{}
}

func NewDaemonThread(node *CFNNode) *DaemonThread {
	return &DaemonThread{
		node:     node,
		interval: time.Second * 5,
	}
}

func (t *DaemonThread) check() {

}

func (t *DaemonThread) run() {
	t.ticker = time.NewTicker(t.interval)
	t.stop = make(chan struct{})

	defer t.ticker.Stop()

	select {
	case <-t.ticker.C:
		t.check()
	case <-t.stop:
		return
	}
}

func (t *DaemonThread) Start() {
	go t.run()
}

func (t *DaemonThread) Stop() {
	close(t.stop)
}
