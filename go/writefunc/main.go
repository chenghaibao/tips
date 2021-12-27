package main

import "fmt"

type option struct {
	Username string
	Password string
	Service  string
}

type clientConn struct {
	timeout int
	service option
}

// 调用接口
type optionInterface interface {
	run(*option)
}

type funcOption struct {
	Function func(*option)
}

// 最后执行参数
func (this *funcOption) run(do *option) {
	this.Function(do)
}

// 参数叠加接口
func withPassword(s string) optionInterface {
	// 定义的是function
	return &funcOption{func(o *option) {
		o.Password = s
	}}
}

// 参数叠加接口
func withUserName(s string) optionInterface {
	// 定义的是function
	return &funcOption{func(o *option) {
		o.Username = s
	}}
}

// 默认参数
func defaultOptions(address string) option {
	return option{
		Service: address,
	}
}

// 初始化参数
func NewClient(address string, opts ...optionInterface) {
	test := &clientConn{
		timeout: 30,
		service: defaultOptions(address),
	}
	//循环调用opts
	for _, opt := range opts {
		opt.run(&test.service)
	}

	fmt.Println(test.service.Service, test.service.Username)
}
func main() {
	NewClient("127.0.0.1", withPassword("asdsa"), withUserName("hb"))
}
