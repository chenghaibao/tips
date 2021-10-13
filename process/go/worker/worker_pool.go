package worker

import (
	"context"
	"sync"
)

type Pool struct {
	ctx    context.Context
	cancel context.CancelFunc
	workerMap sync.Map
	node chan int
}

func (p *Pool) put(workerId string,worker int) {
	p.workerMap.Store(workerId, worker)
}
