package main

import "fmt"

//常量
const Name string = "lisa"

//变量
var age int = 18

//一般类型申明
type gender int

//结构申明
type Dog struct {
	Name string
}

//接口
type IAnimal interface {
	Shout() string
}

//函数
func (g Dog) Shout() string {
	fmt.Println("dog shout wow")
	return "wow~"
}

func main() {
	var animal IAnimal
	var dog = Dog{"大黄"}
	animal = dog
	result := animal.Shout()
	fmt.Println("shout", result)
}
