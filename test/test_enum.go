package main

import "fmt"

//方法一
type DemandType int

const (
	TuiHuo DemandType = iota + 1
	TuiWuTuiHuo
)

func (c DemandType) String() string {
	switch c {
	case TuiHuo:
		return "退货"
	case TuiWuTuiHuo:
		return "退物退货11"
	}
	return "N/A"
}

func (c DemandType) Enum_String() string {
	switch c {
	case TuiHuo:
		return TuiHuo.String() //也以直接赋值
	case TuiWuTuiHuo:
		return "退物退货"
	}
	return "其他"
}

//方法一
var Enum_DemandType = map[int]string{
	1: "退货",
	2: "退物退货",
}

func IntToEnumValue(i int, c map[int]string) string {
	v, is_ok := c[i]
	if !is_ok {
		//弊端就是其他类型，需要自己再处理
		v = "其他"
	}
	return v
}
func StringToEnumValue(i string, c map[string]string) string {
	v, is_ok := c[i]
	if !is_ok {
		//弊端就是其他类型，需要自己再处理
		v = "其他"
	}
	return v
}

func main() {
	input_demand := 1
	//方法一
	fmt.Println(ToEnumValue(input_demand, Enum_DemandType))

	//方法二
	fmt.Println(DemandType(input_demand).Enum_String())
}
