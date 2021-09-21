package main

import (
	"fmt"
)

type SelectFactory interface {
	Go() string
}

type Zoo struct{}

func (z *Zoo) Go() string {
	return "动物圆"
}

type Animal struct{}

func (a *Animal) Go() string {
	return "动物"
}

func main() {
	//简单工厂
	var factory SelectFactory
	AnimalType := "zoo"

	switch AnimalType {
	case "zoo":
		factory = new(Zoo)
	case "animal":
		factory = new(Animal)
	default:
		panic("no such translator")
	}

	fmt.Println(factory.Go())

}
