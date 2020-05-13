package main

import "fmt"

const a = iota
const b = iota

const (
	c = iota
	d = iota
	e //可省略iota
	_ //跳值占位
	f
	g = iota + 4 //非小组内首个表达式，后续不继承该隐式表达式
	h
	i
)

const (
	j, k = iota + 1, iota + 3
	l, m        //i和m分别基础j和k的表达式
	n    = iota //n基础j l的iota的值
)

func main() {
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("c", c)
	fmt.Println("d", d)
	fmt.Println("e", e)
	fmt.Println("f", f)
	fmt.Println("g", g)
	fmt.Println("h", h)
	fmt.Println("i", i)

	fmt.Println("隐式表达式继承")
	fmt.Println("j", j)
	fmt.Println("k", k)
	fmt.Println("l", l)
	fmt.Println("m", m)
	fmt.Println("n", n)
}
