package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"strconv"
)

type Students []Student

type Student struct {
	Name      string
	No        int
	IsMonitor bool
	Courses   []string
	SeatMates []SeatMate
}

type SeatMate struct {
	S_Name string
	Stu_No int
}

func main() {
	//test1()
	test_reflect()
}

func test_excel1()  {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row, row1 *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	row = sheet.AddRow()
	row.SetHeightCM(0.8)
	cell = row.AddCell()
	cell.Value = "姓名"
	cell = row.AddCell()
	cell.Value = "年龄"

	for i := 1; i < 5; i++ {
		row1 = sheet.AddRow()
		row1.SetHeightCM(0.7)
		cell = row1.AddCell()
		cell.Value = "狗子00" + strconv.Itoa(i)
		cell = row1.AddCell()
		cell.Value = strconv.Itoa(18 + i)
	}

	err = file.Save("test_write.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func test_reflect()  {

}