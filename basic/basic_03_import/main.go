package main

import (
	"fmt"
	"go-abc/basic/basic_03_import/animal"
	"go-abc/basic/basic_03_import/life"//重复引用只会有一次

	//. "fmt" 可直接使用fmt包中的函数。但是不建议如此使用，可能导致函数或变量重复
	//_ "fmt" 仅初始化fmt包中的init
	//pt "fmt" 使用别名去使用fmt包中的函数
)

func init()  {
	fmt.Println("main init")
}

func main()  {
	fmt.Println("main start")
	animal.Study()
	fmt.Println("main end")

}

func how_live()  {
	life.Live()
}