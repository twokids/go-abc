package main

import (
	"fmt"
	"math"
)

func main() {
	var age int // 变量声明
	fmt.Println("my age is", age)

	age = 10 //变量赋值
	fmt.Println("my age is", age)

	age = 30 //变量赋值
	fmt.Println("my age is", age)

	var name string = "allen" // 声明变量并初始化
	fmt.Println("my name is", name)

	var num = 29 // 可以推断类型
	fmt.Println("my num is", num)

	var width, height int = 100, 50 // 声明多个变量
	fmt.Println("width is", width, "height is", height)

	var a, b int //省略初始化，默认值为0
	fmt.Println("width is", a, "height is", b)

	//一个语句中声名不同类型变量
	var (
		c = "jie"
		d = 24
		e int
	)
	fmt.Println("my c is", c, ", d is", d, "and e is", e)

	m, n := "allen", 29 // 简短声明
	fmt.Println("m name is", m, "n is", n)

	//变量也可以在运行时进行赋值
	q, p := 145.8, 543.8
	e1 := math.Min(q, p)
	fmt.Println("minimum value is ", e1)

}
