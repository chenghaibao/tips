package main

import (
	"hb_process/aggent"
	"log"
	"runtime"
)

func main() {

	//// 全部核心运行程序
	runtime.GOMAXPROCS(runtime.NumCPU())
	// 日志等级
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	aggent.Execute()

}
