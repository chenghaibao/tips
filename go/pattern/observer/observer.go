package main

import (
	"fmt"
	"reflect"
	"runtime"
)
type Observable interface {
	Attach(observer ...ObserverInterface) Observable
	Detach(observer ObserverInterface) Observable
	Notify() error
}

type ObservableConcrete struct {
	observerList []ObserverInterface
}

type ObserverInterface interface {
	// 自身的业务
	Do(o Observable) error
}

// @param $observer ObserverInterface 观察者列表
func (o *ObservableConcrete) Attach(observer ...ObserverInterface) Observable {
	o.observerList = append(o.observerList, observer...)
	return o
}

// Detach 注销观察者
// @param $observer ObserverInterface 待注销的观察者
func (o *ObservableConcrete) Detach(observer ObserverInterface) Observable {
	if len(o.observerList) == 0 {
		return o
	}
	for k, observerItem := range o.observerList {
		if observer == observerItem {
			fmt.Println(runFuncName(), "注销:", reflect.TypeOf(observer))
			o.observerList = append(o.observerList[:k], o.observerList[k+1:]...)
		}
	}
	return o
}

// Detach 注销观察者
// @param $observer ObserverInterface 待注销的观察者
func (o *ObservableConcrete) Notify() (err error) {
	for _, observer := range o.observerList {
		if err = observer.Do(o); err != nil {
			return err
		}
	}
	return nil
}

type OrderStatus struct {
}

// Do 具体业务
func (observer *OrderStatus) Do(o Observable) (err error) {
	fmt.Println(runFuncName(), "修改订单状态...")
	return
}

// OrderStatusLog 记录订单状态变更日志
type OrderStatusLog struct {
}

// Do 具体业务
func (observer *OrderStatusLog) Do(o Observable) (err error) {
	fmt.Println(runFuncName(), "记录订单状态变更日志...")
	return
}

func main() {
	// 创建 取消发货单 “主题”
	fmt.Println("----------------------- 取消发货单 “主题”")
	deliverBillCancelSubject := &ObservableConcrete{}
	deliverBillCancelSubject.Attach(
		&OrderStatus{},
		&OrderStatusLog{},
	)
	deliverBillCancelSubject.Notify()
}

// 获取正在运行的函数名
func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}
