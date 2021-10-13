package node

import (
	"go.uber.org/zap"
	"log"
	"sync"
)

type PHPAppNodeManage struct {
	*PHPAppNode
	nodeLock sync.Mutex
}

func NewAppNodeManage() {
	agent := NewAppAgent("master")
	go func() {
		err := agent.ListenAndServe()
		if err != nil {
			log.Printf("app node manager exit", zap.Error(err))
		}
	}()
}


func NewAppNode() {
	agent := NewAppAgent("child")
	go func() {
		err := agent.ListenAndServe()
		if err != nil {
			log.Printf("app node child exit", zap.Error(err))
		}
	}()
}