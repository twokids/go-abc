package main

import (
	"fmt"
	"time"
)

//结构申明
type BaseModel struct {
	ID         int
	CreateName string
	Createtime time.Time
}

type User struct {
	BaseModel
	ID       string
	UserName string
}

//函数
func (g BaseModel) InitDate() {
	g.CreateName = "南山"
	g.Createtime = time.Now()
}

func main() {
	test()

	var u User = User{UserName: "lisa", ID: "101"}
	u.InitDate()

	fmt.Println("shout", u)
}

func test() {

	t := time.Now()
	fmt.Println(t.Format("20060102150405"))

	//当前时间戳
	t1 := time.Now().Unix()
	fmt.Println(t1)
	//时间戳转化为具体时间
	fmt.Println(time.Unix(t1, 0).String())

	//基本格式化的时间表示
	fmt.Println(time.Now().String())

	fmt.Println(time.Now().Format("2006year 01month 02day"))
	hour := t.Hour()
	fmt.Print(hour)

	time.Now().Add(time.Hour*-1)

	ti:=time.Now().Before(time.Now().Add(time.Hour*-1))
	fmt.Print(ti)
}
