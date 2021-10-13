package worker

import (
	"context"
	"hb_process/worker/node"
	"time"
)

// 目前 制作数量不足时自动扩容的进程调度，  超时，自动检查，等等没有做进程调度
func Start(ctx context.Context) *Pool {
	pool := &Pool{node: make(chan int, 12)}
	// 一秒中检查一次
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		select {
		case <-ctx.Done():
			return nil
		default:
			if len(pool.node) < 12 {
				node.NewAppNode()
				// 存入map中  2等于 子进程 conn
				pool.put("子进程pid",2)
			}
		}
	}
	return pool
}
