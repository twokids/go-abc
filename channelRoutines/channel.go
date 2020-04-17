package main

import (
	"fmt"
	"strconv"
	"time"
)

/**
在这个示例中，我们的副协程在向channel中写入数据的时候，有一秒的延迟

然后我们发现控制台断断续续的打印出了内容

说明我们从channel中读取消息是个阻塞的行为
*/
func main() {
	/**
	  声明一个长度为 1 的channel
	*/
	messages := make(chan string)

	/**
	  该协程每隔一秒向channel发送一个数据
	*/
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second * 1)
			messages <- "Hello, 这是第 " + strconv.Itoa(i) + " 次传递消息"
		}
	}()

	/**
	  读取channel中的数据，这是个阻塞的操作
	*/
	for i := 0; i < 5; i++ {
		fmt.Println(<-messages)
	}
}
