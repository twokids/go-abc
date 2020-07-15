package main

import "fmt"

type Person struct {
	name string
	sex string
	age int
}

type Student1 struct {
	Person
	mystr
	id int
	addr string
}

// 自定义类型
type mystr string

func main()  {
	s1:=Student1{Person{"jay","women",18},"-",1,"上海"}
	fmt.Println(s1)

	s2:=Student1{Person: Person{name:"lisa",sex:"man"}}
	fmt.Println(s2)

	var s3 Student1

	s3.addr="华侨银行大厦"
	s3.mystr="--"
	s3.name="matt"
	fmt.Println(s3)

	s3.Person.name="amy"
	fmt.Println(s3)

	fmt.Println(s3.Person.name,s3.mystr)

	main2()
}



type User struct {
	id   int
	name string
}

func (self User) Test() {
	fmt.Printf("%p, %v\n", self.id, self.name)
}



func main2() {

	u := User{1, "Tom"}
	mValue := u.Test

	u.id, u.name = 2, "Jack"
	u.Test()

	mValue()


	//u := User{1, "Tom"}
	//u.Test()
	//
	//mValue := u.Test
	//mValue() // 隐式传递 receiver
	//
	//
	//u.id, u.name = 2, "Jack"
	//u.Test()
	//mValue()
	//
	//mExpression := (*User).Test
	//mExpression(&u) // 显式传递 receiver

}