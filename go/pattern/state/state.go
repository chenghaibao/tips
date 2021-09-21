package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

//------------------------------------------------------------
//状态模式
//------------------------------------------------------------

// Context 上下文
type Context struct {
	Tel        string // 手机号
	Text       string // 短信内容
	TemplateID string // 短信模板ID
}

type SmsServiceInterface interface {
	Send(ctx *Context) error
}

// ServiceProviderAliyun 阿里云
type ServiceProviderAliyun struct {
}

// Send 具体的发送短信逻辑
func (s *ServiceProviderAliyun) Send(ctx *Context) error {
	fmt.Println(runFuncName(), "【阿里云】短信发送成功，手机号:"+ctx.Tel)
	return nil
}

// ServiceProviderTencent 腾讯云
type ServiceProviderTencent struct {
}

// Send 具体的发送短信逻辑
func (s *ServiceProviderTencent) Send(ctx *Context) error {
	fmt.Println(runFuncName(), "【腾讯云】短信发送成功，手机号:"+ctx.Tel)
	return nil
}

type ProviderType string

const (
	// ProviderTypeAliyun 阿里云
	ProviderTypeAliyun ProviderType = "aliyun"
	// ProviderTypeTencent 腾讯云
	ProviderTypeTencent ProviderType = "tencent"
)

var (
	// stateManagerInstance 当前使用的服务提供商实例
	stateManagerInstance *StateManager
)

type StateManager struct{
	currentProviderType ProviderType
	currentProvider SmsServiceInterface
	// 更新状态时间间隔
	setStateDuration time.Duration
}
// initState 初始化状态
func (m *StateManager) initState(duration time.Duration){
	// 初始化
	m.setStateDuration = duration
	m.setState(time.Now())
	// 定时器更新状态
	go func() {
		for {
			select{
				case t := <-time.NewTicker(m.setStateDuration).C:
					m.setState(t)
			}
		}
	}()
}

// setState 设置状态
// 根据短信云商回调的短信发送成功率 得到下阶段发送短信使用哪个厂商的服务
func (m *StateManager) setState(t time.Time) {
	// 这里用随机模拟
	ProviderTypeArray := [2]ProviderType{
		ProviderTypeAliyun,
		ProviderTypeTencent,
	}
	m.currentProviderType = ProviderTypeArray[rand.Intn(len(ProviderTypeArray))]

	switch m.currentProviderType {
	case ProviderTypeAliyun:
		m.currentProvider = &ServiceProviderAliyun{}
	case ProviderTypeTencent:
		m.currentProvider = &ServiceProviderTencent{}
	default:
		panic("无效的短信服务商")
	}
	fmt.Printf("时间：%s| 变更短信发送厂商为: %s \n", t.Format("2006-01-02 15:04:05"), m.currentProviderType)
}

func(m *StateManager) getState() SmsServiceInterface{
	return m.currentProvider
}

// GetState 获取当前状态
func GetState() SmsServiceInterface {
	return stateManagerInstance.getState()
}

// 获取正在运行的函数名
func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}
func main() {
	// 初始化状态管理
	stateManagerInstance = &StateManager{}
	stateManagerInstance.initState(300 * time.Millisecond)

	// 模拟发送短信的接口
	sendSms := func() {
		// 发送短信
		GetState().Send(&Context{
			Tel:        "+8613666666666",
			Text:       "3232",
			TemplateID: "TYSHK_01",
		})
	}
	sendSms()
	time.Sleep(1 * time.Second)
	sendSms()
	time.Sleep(1 * time.Second)
	sendSms()
}
