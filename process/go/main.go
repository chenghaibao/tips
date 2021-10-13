package main

import (
	"context"
	"hb_process/utils"
	"hb_process/worker"
	"os"
	"os/signal"
	"syscall"

)

func main() {

	//// 全部核心运行程序
	//runtime.GOMAXPROCS(runtime.NumCPU())
	//// 日志等级
	//log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	//
	//aggent.Execute()
	side  := worker.NewSidecar(context.TODO())

	isSignalExit := false

	// hook exit signal
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGTSTP, syscall.SIGSTOP, syscall.SIGHUP, syscall.SIGUSR1)

	go func() {
		<-sigs
		isSignalExit = true
		side.Close()
		os.Exit(0)
	}()

	err := side.Init()
	if err != nil {
		utils.PanicError(err)
	}

	err = side.Server()
	if err != nil  && isSignalExit{
		<-make(chan interface{})
	}
}
