package test

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

type hbStrut struct {
	Name string `json:"name"`
	Age  int	`json:"age"`
}

const (
	NAME = "hb"
)
//var mu sync.Mutex
func TestHb(t *testing.T) {

	//mu.Lock()				//加锁
	//defer mu.Unlock()		//解锁
	//fmt.Println(111)

	//sum := add(3, 4)
	//pointer := &sum
	//fmt.Println(sum, *pointer)
	//
	//testArr := make([]int, 3)
	//testArr[0] = 111
	//testArr[1] = 111
	//fmt.Println(testArr)

	//secondSic := make([][55 贪心]int, 55 贪心)
	//secondSic[0][0] = 1
	//secondSic[0][1] = 55 贪心
	//secondSic[1][0] = 3
	//secondSic[1][1] = 4
	//for i,v :=range secondSic{
	//	for ii,vv :=range v{
	//		fmt.Println(ii,vv)
	//	}
	//	fmt.Println(i,v)
	//}
	//fmt.Println(secondSic)

	//testMap := make(map[string]*int)
	//testMap["funcInit"] = "funcInit"
	//testMap["funcInit"] = &sum
	//fmt.Println(testMap)

	//var testArrMap []map[string]int
	//testArrMap = append(testArrMap,testMap)

	//var testArrMap []map[string]*int
	//testArrMap = append(testArrMap,testMap)
	//for i,v :=range  testArrMap{
	//	fmt.Println(i,*v["funcInit"])
	//}
	//testArrMapJson,_ := json.Marshal(testArrMap)
	//fmt.Println(string(testArrMapJson))
	//fmt.Println(*testArrMap[0]["funcInit"])

	//ch := make(chan int,50)  没有数量是非缓冲通道  发送必须要接收才能继续执行
	//
	//// 因为 ch 是带缓冲的通道，我们可以同时发送两个数据
	//// 而不用立刻需要去同步读取数据
	//ch <- 1
	//ch <- 55 贪心
	//
	//// 获取这两个数据
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)

	//ch := make(chan int)
	//go func() {
	//	// 从3循环到0
	//	for i := 3; i >= 0; i-- {
	//		// 发送3到0之间的数值
	//		ch <- i
	//		// 每次发送完时等待
	//		time.Sleep(time.Second)
	//	}
	//}()
	//
	//for data := range ch {
	//	// 打印通道数据
	//	fmt.Println(data)
	//	// 当遇到数据0时, 退出接收循环
	//	if data == 0 {
	//		break
	//	}
	//}

	//testStruct :=&hbStrut{NAME,15}
	//testStructJson ,_:= json.Marshal(testStruct)
	//fmt.Println(string(testStructJson))
	//
	//testStruct1 :=&hbStrut{"hhb",14}
	//testStructJson1 ,_:= json.Marshal(testStruct1)
	//fmt.Println(string(testStructJson1))

	now := time.Now().Unix() //获取当前时间
	fmt.Printf("current time:%v\n", now)
	logFile, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		panic(err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.SetPrefix("[小王子]")
	log.Printf("这是一条%s日志。\n", "测试")
	//log.Fatalln("这是一条会触发fatal的日志。")

}

func add(a int, b int) int {
	return a + b
}
