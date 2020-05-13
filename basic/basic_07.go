package main

import (
	"fmt"
)

func suanshu() {
	a, b := 2, 3
	fmt.Println("a+b", a+b)
	fmt.Println("a-b", a-b)
	fmt.Println("a*b", a*b)
	fmt.Println("a/b", a/b)
	fmt.Println("a%b", a%b)
	fmt.Print("a++")
	a++
	fmt.Println(a)
	c := b + 1
	b--
	fmt.Println("c", c)
	fmt.Println("b--")
	fmt.Println(b)
}

func guanxi() {
	a, b := 2, 3
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("a==b", a == b)
	fmt.Println("a>b", a > b)
	fmt.Println("a>=b", a >= b)
	fmt.Println("a<b", a < b)
	fmt.Println("a<=b", a <= b)
	fmt.Println("a!=b", a != b)
}
func luoji() {

	a, b, c := true, true, false
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("c", c)
	fmt.Println("a&&b", a && b)
	fmt.Println("a&&c", a && c)
	fmt.Println("a||b", a || b)
	fmt.Println("a||c", a || c)
	fmt.Println("!a", !a)
	fmt.Println("!c", !c)
}

func anwei() {
	/*
		&按位与，都是1则为1
		|按位或，有一个1则为1
		^按位异或，对应位的数值不同则为1
		<<左移，把整数体向左移动
		>>右移，把整数体按右移动
	*/

	a, b := 3, 4
	//a: 0000 0011
	//b: 0000 0100
	fmt.Println("a,b", a, b)
	fmt.Println("a&b", a&b)//0
	fmt.Println("a|b", a|b)//7
	fmt.Println("a^b", a^b)//7
	a = a << 2
	fmt.Println("a<<2", a)//12
	b = b >> 1
	fmt.Println("b>>1", b)//2
	fmt.Println("b>>2", b>>2)//0

}

func fuzhi()  {
	a,b:=1,2
	c:=a+b
	fmt.Println("a+b",c)

	c+=a
	fmt.Println("a+b",c)
	fmt.Println("a+b",c)
	fmt.Println("a+b",c)
	fmt.Println("a+b",c)
	fmt.Println("a+b",c)
	fmt.Println("a+b",c)
}

func main() {
	suanshu()
	guanxi()
	luoji()
	anwei()

}
