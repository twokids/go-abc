package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type nanshan int32 // 类型别名

func main() {

	var a int = 1
	var b int32 = 2
	var c float32 = 3
	var d float64 = 4

	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))
	fmt.Println(unsafe.Sizeof(c))
	fmt.Println(unsafe.Sizeof(d))

	//var e bool =1 必须true or false
	//fmt.Println(e)

	var f int32 = 2
	var g float64 = 3
	var h bool
	var i string
	var j complex64

	fmt.Println("f", f)
	fmt.Println("g", g)
	fmt.Println("h", h)
	fmt.Println("i", i)
	fmt.Println("j", j)

	var k nanshan = 4
	var l nanshan = 5
	fmt.Println("\n 类型别名 k", k)
	fmt.Println(reflect.TypeOf(k))
	fmt.Println(unsafe.Sizeof(k))

	//fmt.Println(f+k)不同类型不可相加
	//fmt.Println(g+k)

	fmt.Println("k+l", k+l)
}
