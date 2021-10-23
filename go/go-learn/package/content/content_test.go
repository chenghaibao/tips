package content

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestContentCancel(t *testing.T){
	ctx, cancel := context.WithCancel(context.Background())
	go ctxCancel(ctx)

	//10秒后取消ctxCancel
	time.Sleep(10 * time.Second)
	cancel()
}

func ctxCancel(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			log.Printf("done")
			return
		default:
			log.Printf("work")
		}
	}
}

func TestContentBackground(t *testing.T){
	ctx, cancel := context.WithCancel(context.Background())
	go ctxBackground(ctx)

	//10秒后取消doStuff
	time.Sleep(5 * time.Second)
	cancel()
}

//每1秒work一下，同时会判断ctx是否被取消，如果是就退出
func ctxBackground(ctx context.Context) {
	var i = 1
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("done")
			return
		default:
			fmt.Printf("work %d seconds: \n", i)
		}
		i++
	}
}


func TestContentDeadline(t *testing.T){

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	// go deadline(ctx)
	go deadline(ctx)
	time.Sleep(5 * time.Second)
	cancel()
}

//每1秒work一下，同时会判断ctx是否被取消，如果是就退出
func deadline(ctx context.Context) {
	var i = 1
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("done")
			return
		default:
			fmt.Printf("work %d seconds: \n", i)
		}
		i++
	}
}


func TestContentTimeOut(t *testing.T) {
	// 创建继承Background的子节点Context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	go Timeout(ctx)

	//模拟程序运行 - Sleep 10秒
	time.Sleep(5 * time.Second)
	cancel() // 3秒后将提前取消 doSth goroutine
}

//每1秒work一下，同时会判断ctx是否被取消，如果是就退出
func Timeout(ctx context.Context) {
	var i = 1
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("done")
			return
		default:
			fmt.Printf("work %d seconds: \n", i)
		}
		i++
	}
}