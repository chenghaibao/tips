package main

import (
	"fmt"
	"hb_responsibitity/responsefunc"
)

func checkFuncAuth() (string, error) {
	return "auth", nil
}

func checkFuncParams() error {
	fmt.Println("params")
	return nil
}

func checkFuncDefault() error {
	fmt.Println("default")
	return nil
}


func checkFuncAdd() error {
	fmt.Println("add")
	return nil
}

func main() {
	// 普通方法策略模式
	link := responsefunc.NewLink(checkFuncAuth)
	link.Whens(checkFuncParams, "bbb", "ccc").
		Default(checkFuncDefault)

	//fmt.Println(link.Run())

	// 中间方法策略模式
	_ = responsefunc.NewPipeline().
		Add(checkFuncAdd).
		AddLink(link).
		Run()

}
