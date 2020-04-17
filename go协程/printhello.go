package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("print hello started")
	//time.Sleep(10 * time.Millisecond)

	//开启另一个printhello携程和main携程并行，但是main如果执行完了。会关闭所有程序。也把printhello关闭掉
	go printHello()


	time.Sleep(10 * time.Millisecond)
	fmt.Println("print hello end")
}

func printHello() {
	time.Sleep(9 * time.Millisecond)
	fmt.Println("hello everyone")
}
