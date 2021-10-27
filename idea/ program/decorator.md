
# 修饰器

## 修饰 解决哪些问题

1. 重复方法的使用
2. 流程的清晰

## 示例代码
https://coolshell.cn/articles/17929.html

### 用法 1
```
// 用法 1
package main
import "fmt"

func decorator(f func(s string)) func(s string) {
	return func(s string) {
		fmt.Println("Started")
		f(s)
		fmt.Println("Done")
	}
}

func Hello(s string) {
	fmt.Println(s)
}

func main() {
	decorator(Hello)("Hello, World!")
	hello := decorator(Hello)
	hello("Hello")
}
```

### 用法 2
```
// 用法 2
type SumFunc func(s string) string

func decorator(f SumFunc) SumFunc {
	return func(s string) string {
		fmt.Println("start")
		return f(s)
	}
}

func Hello(s string) string {
	fmt.Println(s)
	return s
}

func main() {
	hello := decorator(Hello)
	hello("Hello")
}

```
