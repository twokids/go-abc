package main

import "fmt"

const apple_num int = 10
const orange_num = 9

const (
	banana_num      int = 13
	watermelons_num     = 15 //西瓜
)

//芒果，椰子
const mango_num, coconut_num int = 3, 5
const strawberry_num, pear_num = 7, 8

func main() {
	fmt.Println("apple number ", apple_num)
	fmt.Println("orange number ", orange_num)

	fmt.Println("banana num ", banana_num)
	fmt.Println("watermelons num ", watermelons_num)

	fmt.Println("mango_num ", mango_num)
	fmt.Println("coconut_num ", coconut_num)

	fmt.Println("strawberry num ", strawberry_num)
	fmt.Println("pear num ", pear_num)

}
