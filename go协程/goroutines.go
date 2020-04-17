package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func getChars(s string) {
	for _, c := range s {
		fmt.Printf("%c at time %v\n", c, time.Since(start))
		time.Sleep(10 * time.Millisecond)
	}
}

func getDigits(s []int) {
	for _, d := range s {
		fmt.Printf("%d at time %v\n", d, time.Since(start))
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	fmt.Println("exec started at time", time.Since(start))
	go getChars("hello")
	go getDigits([]int{1, 2, 3, 4, 5})
	time.Sleep(200 * time.Millisecond)
	//理论上主协程会休眠 200ms，因此其他 Go 协程要赶在主协程唤醒之前做完自己的工作，因为主协程唤醒之后就会导致程序退出。
	//getChars 协程每打印一个字符就会休眠 10ms，之后控制权就会传给 getDigits 协程，getDigits 协程每打印一个数字后就休眠 30ms，
	//若 getChars 协程唤醒，则会把控制权传回 getChars 协程，如此往复。

	//但是使用 time.Sleep 只是一个让我们获取理想结果的一个小技巧。
	//在实际生产环境中，我们无法知晓一个 Go 协程到底需要执行多长的时间，
	//因而在 main 函数里面添加一个 time.Sleep 并不是一个解决问题的方法。
	//我们希望 Go 协程在执行完毕后告知主协程运行的结果。
	//在目前阶段，我们还不知道如何向其他 Go 协程传递以及获取数据，简而言之，就是与其他 Go 协程进行通信。
	//这就是 channels 引入的原因。
	fmt.Println("\nexec end at time", time.Since(start))
}
