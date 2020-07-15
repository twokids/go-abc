package main

import "fmt"

func main()  {
	//var a1=make([]int,5,8)
	//for i:=0;i<10;i++{
	//	a1=append(a1,i)
	//}
	//
	//fmt.Println(a1)
	//
	//a2:=[]int{1,2,3,4,5,6}
	//res := append(a2[:1], a2[2:]...)
	//fmt.Println(res)
	//
	//
	//u := User{1, "Tom"}
	//fmt.Printf("User: %p, %v\n", &u, u)
	//
	//mv := User.TestValue
	//mv(u)
	//
	//mp := (*User).TestPointer
	//mp(&u)
	//
	//mp2 := (*User).TestValue // *User 方法集包含 TestValue。签名变为 func TestValue(self *User)。实际依然是 receiver value copy。
	//mp2(&u)


	//a11 := []int{1, 2, 3}
	//a12 := a11[:1]
	//a12 = append(a12, 99)
	//a12 = append(a12, 99)
	//a11[1] = 100
	//fmt.Println(a11) // 1 99 99 99
	//fmt.Println(a12) //1 99 99 99


	var str = [5]string{  4: "tom", 1: "to2m","hello world"}
	fmt.Println(str[0])

	d := [...]struct {
		name string
		age  uint8
			id int

	}{
		{"user1", 10,2}, // 可省略元素类型。
		{"user2", 20,5}, // 别忘了最后一行的逗号。
	}

	fmt.Println(d)
}
