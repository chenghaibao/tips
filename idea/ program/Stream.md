
# 流式技术

## 解决问题

1. 代码调用链
2. 代码重构

## 示例代码
```
package _program

import (
	"errors"
	"fmt"
	"testing"
)

type Cmd func(interface{}) (interface{}, error)

//任务节点结构定义
type Stream struct {
	//任务链表首节点,其他非首节点此指针永远指向首节点
	firstStream *Stream
	//任务链表下一个节点，为空表示任务结束
	nextStream *Stream
	//当前任务对应的执行处理函数，首节点没有可执行任务，处理函数指针为空
	cmd Cmd
}

/**
创建新的流
**/
func NewStream() *Stream {
	//生成新的节点
	stream := &Stream{}
	//设置第一个首节点，为自己
	//其他节点会调用run方法将从firs指针开始执行，直到next为空
	stream.firstStream = stream
	//fmt.Println("new first", stream)
	return stream
}

/**
流结束
arg为流初始参数，初始参数放在End方法中是考虑到初始参数不需在任务链中传递
**/
func (this *Stream) Go(arg interface{}) (interface{}, error) {
	//设置为任务链结束
	this.nextStream = nil
	//fmt.Println("first=", this.firstStream, "second=", this.firstStream.nextStream)
	//检查是否有任务节点存在，存在则调用run方法
	//run方法是首先执行本任务回调函数指针，然后查找下一个任务节点，并调用run方法
	if this.firstStream.nextStream != nil {
		return this.firstStream.nextStream.run(arg)
	} else {
		//流式任务终止
		return nil, errors.New("Not found execute node.")
	}
}
func (this *Stream) run(arg interface{}) (interface{}, error) {
	//fmt.Println("run,args=", args)
	//执行本节点函数指针
	result, err := this.cmd(arg)
	//然后调用下一个节点的Run方法
	if this.nextStream != nil && err == nil {
		return this.nextStream.run(result)
	} else {
		//任务链终端，流式任务执行完毕
		return result, err
	}
}
func (this *Stream) Next(cmd Cmd) *Stream {
	//创建新的Stream，将新的任务节点Stream连接在后面
	this.nextStream = &Stream{}
	//设置流式任务链的首节点
	this.nextStream.firstStream = this.firstStream
	//设置本任务的回调函数指针
	this.nextStream.cmd = cmd

	return this.nextStream
}

//起床
func GetUP(arg interface{}) (interface{}, error) {
	t, _ := arg.(string)
	fmt.Println("铃铃.......", t, "###到时间啦，再不起又要迟到了！")
	return "醒着的状态", nil
}

//蹲坑
func GetPit(arg interface{}) (interface{}, error) {
	s, _ := arg.(string)
	fmt.Println(s, "###每早必做的功课，蹲坑！")
	return "舒服啦", nil
}

//洗脸
func GetFace(arg interface{}) (interface{}, error) {
	s, _ := arg.(string)
	fmt.Println(s, "###洗脸很重要！")
	return "脸已经洗干净了，可以去见人了", nil
}

//刷牙
func GetTooth(arg interface{}) (interface{}, error) {
	s, _ := arg.(string)
	fmt.Println(s, "###刷牙也很重要！")
	return "牙也刷干净了，可以放心的大笑", nil
}

//吃早饭
func GetEat(arg interface{}) (interface{}, error) {
	s, _ := arg.(string)
	fmt.Println(s, "###吃饭是必须的(需求变更了，原来的流程里没有，这次加上)")
	return "吃饱饱了", nil
}

//换衣服
func GetCloth(arg interface{}) (interface{}, error) {
	s, _ := arg.(string)
	fmt.Println(s, "###还要增加一个换衣服的流程！")
	return "找到心仪的衣服了", nil
}

//出门
func GetOut(arg interface{}) (interface{}, error) {
	s, _ := arg.(string)
	fmt.Println(s, "###一切就绪，可以出门啦！")
	return "", nil

}

func TestAdd(t *testing.T) {
	NewStream().
		Next(GetUP).
		Next(GetPit).
		Next(GetTooth).
		Next(GetFace).
		Next(GetEat).//需求变更了后加上的
		Next(GetCloth).
		Next(GetOut).
		Go("2018年1月28日8点10分")
}
```
