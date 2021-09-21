package responsefunc

import "fmt"

// 接口约束
type Func func() error
type ErrorFunc func(err error) error
type WhenFunc func() (string, error)

type Link struct {
	method      WhenFunc
	defaultFunc Func
	error       ErrorFunc
	children    []*LinkChild
	parent      *Link
}

type LinkChild struct {
	conditions []string
	method     Func
}

func NewLink(method WhenFunc) *Link {
	return &Link{method: method}
}

func (l *Link) Whens(method Func, conditions ...string) *Link {
	l.children = append(l.children, &LinkChild{method: method, conditions: conditions})
	return l
}

func (l *Link) Default(method Func) *Link {
	l.defaultFunc = method
	return l
}

func (l *Link) Catch(method ErrorFunc) *Link {
	l.error = method
	return l
}


func (l *Link) Run() error {
	// 执行
	result, err := l.method()
	// 定义
	nothing := true
	// 遍历循环
	for _, link := range l.children {
		// 判断是否存在条件
			// 执行
			nothing = false
			// 执行
			err = link.method()

			fmt.Println(result)
	}
	// 若未执行任何符合条件的func
	if nothing {
		err = l.defaultFunc()
	}
	// 执行错误处理
	if l.error != nil {
		err = l.error(err)
	}
	// 返回
	return err
}


