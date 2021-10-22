package string

import (
	"fmt"
	"strings"
	"testing"
)

type Reader struct {
	// contains filtered or unexported fields
	str *strings.Reader
}

func TestString(t *testing.T) {
	fmt.Println(strings.Compare("a", "b"))
	// 判断 substr 是否包含在 s 之内。
	fmt.Println(strings.Contains("seafood", "foo"))
	// ields 将字符串 s 分割为一个或多个连续的空格字符的每个实例
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
	// 索引返回 s 中第一个 substr 实例的索引，如果 substr 不存在于 s 中，则返回-1。
	fmt.Println(strings.Index("chicken", "ken"))
	// IndexFunc 将索引返回到满足 f(c) 的第一个
	fmt.Println(strings.IndexByte("golang", 'g'))
	//php implode
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", "))
	// LastIndex 返回 s 中最后一个 substr 实例的索引，如果 substr 不存在于 s 中，则返回-1
	fmt.Println(strings.LastIndex("go gopher", "go"))
	// 字符串替换
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	// explode
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
    // trim
	fmt.Printf("[%q]\n", strings.Trim(" !!! Achtung! Achtung! !!! ", "! "))

	// string Reader
	aa := strings.NewReader("6666")
	st := &Reader{str:aa}
	fmt.Println(st.str.Len())
	fmt.Println(st.str.ReadRune())
}