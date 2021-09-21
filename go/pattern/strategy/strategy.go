package main

import "fmt"

const (
	WPay = "Wpay"
	APay = "APay"
)

// UserContext 上下文
type UserContext struct {
	// 用户选择的支付方式
	PayType string `json:"pay_type"`
}

type Select interface {
	Pay(c *UserContext) (string, error)
}

type Wpay struct {
}

func (w Wpay) Pay(c *UserContext) (string, error) {
	return "wpay", nil
}

type Apay struct {
}

func (w Apay) Pay(c *UserContext) (string, error) {
	return "apay", nil
}

func main() {
	pay := &UserContext{PayType: "APay"}
	var instance Select
	switch pay.PayType {
	case WPay:
		instance = &Wpay{}
	case APay:
		instance = &Apay{}
	default:
		panic("无效的支付方式")
	}
	fmt.Println(instance.Pay(pay))
}
