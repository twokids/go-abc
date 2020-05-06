package main

import (
	"fmt"
	"go-abc/basic/basic_03_import/animal"
	"go-abc/basic/basic_03_import/life"//重复引用只会有一次
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