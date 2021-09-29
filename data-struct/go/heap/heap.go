package main

import (
	"errors"
	"fmt"
)

type Stack []interface{}

func (s *Stack) Push(x interface{}) {
	*s = append(*s, x)
}

func (s *Stack) Pop() (interface{}, error) {
	if len(*s) == 0 {
		return nil, errors.New("slice为空!")
	}
	result := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return result, nil
}

func (s *Stack) Len() int {
	return len(*s)
}

func main() {
	s := new(Stack)
	s.Push(1)
	s.Push(2)
	fmt.Println(s, s.Len())
	s.Pop()
	fmt.Println(s, s.Len())
}
