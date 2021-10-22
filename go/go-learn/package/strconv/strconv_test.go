package strconv

import (
	"fmt"
	"strconv"
	"testing"
)

type strc struct {
}

func TestStr(t *testing.T) {
	s := "2147483647" // biggest int32
	i64, _ := strconv.ParseInt(s, 10, 32)
	fmt.Println(i64)

	s1 := strconv.Quote(`"Fran & Freddie's Diner	☺"`)
	fmt.Println(s1)

	s2 := strconv.QuoteRune('☺')
	fmt.Println(s2)

	c := strconv.IsPrint('\u263a')
	fmt.Println(c)

	v := uint64(42)
	s10 := strconv.FormatUint(v, 10)
	fmt.Printf("%T, %v\n", s10, s10)

}