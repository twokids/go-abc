package main

import (
	"fmt"
	"time"
)

func main() {
	var input interface{}
	input = "muke"

	switch input.(type) {
	case int:
		fmt.Println("这是个input类型")
	case string:
		fmt.Println("这是个字符串类型")
	default:
		fmt.Println("其他类型")
	}

	for i := 0; i < 10; i++ {
		fmt.Println("i,", i)
	}

	fruit := []string{"apple", "banana"}

	for _, value := range fruit {
		fmt.Println("fruit is ", value)
	}
	//死循环
	//One:
	//	fmt.Println("this is one")
	//	time.Sleep(1 * time.Second)
	//
	//	goto One

	goto Two
	fmt.Println("i'm ready to two")
Two:
	fmt.Println("this is two")
	time.Sleep(1 * time.Second)

	for {
		fmt.Println("i ready break")
		break
	}

	for i:=0;i<10;i++{
		if i>5{
			fmt.Println("i比5大",i)
			continue
		}
		fmt.Println("i的值",i)
	}
}
