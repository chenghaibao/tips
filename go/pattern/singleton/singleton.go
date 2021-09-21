package main

import (
	"fmt"
	"sync"
)

//------------------------------------------------------------
//单例模式
//------------------------------------------------------------
//var instance *singleton
//
//type singleton struct {
//	Name string
//	Age  int
//}

// // ---------线程不安全的单例 ---------
//func GetInstance() *singleton {
//	if instance == nil {
//		instance = &singleton{Name: "hb", Age: 24}
//	}
//	return instance
//}

// // ---------线程安全的单例-加锁---------初始化加锁 init
//func GetInstance() *singleton {
//	if instance != nil {
//		return instance
//	}
//
//	mock := sync.Mutex{}
//	mock.Lock()
//	instance = &singleton{Name: "hb", Age: 24}
//	defer mock.Unlock()
//
//	return instance

//}

// // ---------线程安全的单例-原子操作判断实例为空时再加锁---------
//var instance *singleton
//var flag *uint32
//
//type singleton struct {
//	Name string
//	Age  int
//}
//
//func GetInstance() *singleton {
//	if atomic.AddUint32(flag, 0) != 0 {
//		return instance
//	}
//
//	//创建单列加锁
//	mutex := &sync.Mutex{}
//	mutex.Lock()
//	defer mutex.Unlock()
//
//	if instance == nil {
//		instance = &singleton{Name: "hb", Age: 24}
//		// 使用原子操作标示 atomic原子库
//		atomic.AddUint32(flag, 1)
//	}
//
//	return instance
//}

// ---------线程安全的单例-使用sync.Once---------
var instance *singleton
var oc *sync.Once

func init() {
	oc = &sync.Once{}
}

type singleton struct {
	Name string
	Age  int
}

func GetInstance(name string) *singleton {
	oc.Do(func() {
		instance = &singleton{Name: name, Age: 24}
	})
	return instance
}

func main() {
	//线程不安全的单例
	//GetInstance()
	fmt.Println(GetInstance("hb"))
	fmt.Println(GetInstance("b"))
}
