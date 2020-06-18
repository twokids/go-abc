package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"os"
	"path/filepath"
	"strconv"
	"time"
)


type Stu struct {
	Id int
	Name string
}
var list3 = []Stu{{1,"a"},{4,"b"}, {3, "c"}}

var list4 = []Stu{}


func main() {





}

//多sheet excel
func test1() {
	f := excelize.NewFile()
	for i := 0; i < 3; i++ {
		test11(f, i)
	}
	// 根据指定路径保存文件

	f.SetActiveSheet(1)
	outFileName := "book.xlsx"
	filename := ("./files/export/" + outFileName + time.Now().Format("-20060102150405"))
	err := f.SaveAs(filename)
	if err != nil {
		fmt.Println(err)
	}
}
func test11(f *excelize.File, ii int) {

	data11 := []map[string]interface{}{
		map[string]interface{}{"name": "jack", "age": 18, "createtime": "2020/06/01 12:00:00", "stuno": 1001,},
		map[string]interface{}{"name": "mary", "age": 28, "stuno": 1002, "createtime": "2020/06/01 12:00:00",},
		map[string]interface{}{"name": "lisa", "age": 38, "createtime": "2020/06/01 12:00:00",},
	}

	for k, v := range data11 {
		fmt.Println(k)
		fmt.Println(v)
	}

	//// 创建一个工作表
	if ii == 0 {

		f.SetSheetName("Sheet1", "第一个sheet")

	} else {
		f.NewSheet("另一个sheet2")
	}

	//// 设置单元格的值
	//f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// 设置工作簿的默认工作表
	//f.SetActiveSheet(index)

}

//图标类型excel
func test2() {
	categories := map[string]string{"A2": "Small", "A3": "Normal", "A4": "Large", "B1": "Apple", "C1": "Orange", "D1": "Pear"}
	values := map[string]int{"B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8}
	f := excelize.NewFile()
	for k, v := range categories {
		f.SetCellValue("Sheet1", k, v)
	}
	for k, v := range values {
		f.SetCellValue("Sheet1", k, v)
	}
	err := f.AddChart("Sheet1", "E1", `{"type":"col3DClustered","series":[{"name":"Sheet1!$A$2","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$2:$D$2"},{"name":"Sheet1!$A$3","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$3:$D$3"},{"name":"Sheet1!$A$4","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$4:$D$4"}],"title":{"name":"Fruit 3D Clustered Column Chart"}}`)
	if err != nil {
		fmt.Println(err)
	}
	// 根据指定路径保存文件
	err = f.SaveAs("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

func test3() {
	xlsx := excelize.NewFile()

	index := xlsx.NewSheet("Sheet1")
	xlsx.SetCellValue("Sheet1", "A1", "姓名")
	xlsx.SetCellValue("Sheet1", "B1", "年龄")
	xlsx.SetCellValue("Sheet1", "A2", "狗子")
	xlsx.SetCellValue("Sheet1", "B2", "18")

	// Set active sheet of the workbook.
	xlsx.SetActiveSheet(index)
	// Save xlsx file by the given path.
	err := xlsx.SaveAs("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

func test4() {
	// 列标题
	titles := []string{
		"姓名", "年龄",
	}
	// 数据源
	data := []map[string]interface{}{
		map[string]interface{}{"name": "jack", "age": 18,},
		map[string]interface{}{"name": "mary", "age": 28,},
	}

	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")

	for clumnNum, v := range titles {
		sheetPosition := Div(clumnNum+1) + "1"
		fmt.Print(sheetPosition)
		f.SetCellValue("Sheet1", sheetPosition, v)
	}

	//lineNo:=1
	for lineNum, v := range data {
		// Set value of a cell.
		clumnNum := 0
		for _, vv := range v {
			clumnNum++
			sheetPosition := Div(clumnNum) + strconv.Itoa(lineNum+2)
			switch vv.(type) {
			case string:
				f.SetCellValue("Sheet1", sheetPosition, vv.(string))
				break
			case int:
				f.SetCellValue("Sheet1", sheetPosition, vv.(int))
				break
			case float64:
				f.SetCellValue("Sheet1", sheetPosition, vv.(float64))
				break
			default:
				//【】 ++
				break

			}
		}
	}
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save xlsx file by the given path.
	if err := f.SaveAs("Book2.xlsx"); err != nil {
		println(err.Error())
	}
}

type ExcelModel struct {
	FileName    string       //存储文件名
	FilePath    string       //存储路径
	SheetModels []SheetModel //多个sheets
}

type SheetModel struct {
	CurSheetName string                   //当前sheet的名字
	Titles       []ETModel                //sheet的标题
	Datas        []map[string]interface{} //数据
}

type ETModel struct {
	TitleName string
	DataName  string
	Sequence  int
}

func test5() {
	sheet1_data := []map[string]interface{}{
		map[string]interface{}{"name": "jack", "age": 18, "createtime": "2020/06/01 12:00:00", "stuno": 1001,},
		map[string]interface{}{"name": "mary", "age": 28, "stuno": 1002, "createtime": "2020/06/01 12:00:00",},
		map[string]interface{}{"name": "lisa", "age": 38, "createtime": "2020/06/01 12:00:00",},
	}
	sheet1 := SheetModel{
		CurSheetName: "01班学生",
		Titles: []ETModel{
			{TitleName: "姓名", DataName: "name", Sequence: 1,},
			{TitleName: "年龄", DataName: "age", Sequence: 2,},
			{TitleName: "学号", DataName: "stuno", Sequence: 3,},
		},
		Datas: sheet1_data,
	}

	sheet2_data := []map[string]interface{}{
		map[string]interface{}{"name": "张三", "age": 18, "createtime": "2020/06/01 12:00:00", "stuno": 1001,},
		map[string]interface{}{"name": "李四", "age": 28, "stuno": 1002, "createtime": "2020/06/01 12:00:00",},
		map[string]interface{}{"name": "王铁柱", "age": 38, "createtime": "2020/06/01 12:00:00",},
	}
	sheet2 := SheetModel{
		CurSheetName: "02班学生",
		Titles: []ETModel{
			{TitleName: "姓名", DataName: "name", Sequence: 2,},
			{TitleName: "年龄", DataName: "age", Sequence: 3,},
			{TitleName: "学号", DataName: "stuno", Sequence: 1,},
		},
		Datas: sheet2_data,
	}
	ops := ExcelModel{
		FileName:    "2008届大学生入学信息表.xlsx",
		FilePath:    "./file/student/08/",
		SheetModels: []SheetModel{sheet1, sheet2},
	}
	test51(ops)
}

func test51(ops ExcelModel) {
	//创建个文件
	f := excelize.NewFile()
	for _, v := range ops.SheetModels {
		test_sheet(f, v)
	}
	// Save xlsx file by the given path.
	CreateDateDir(ops.FilePath)
	path := ops.FilePath + ops.FileName
	if err := f.SaveAs(path); err != nil {
		println(err.Error())
	}
}

//basePath是固定目录路径
func CreateDateDir(basePath string) (dirPath, dataString string) {
	folderName := time.Now().Format("2006-01-02")
	folderPath := filepath.Join(basePath, folderName)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步
		// 先创建文件夹A
		os.Mkdir(folderPath, 0777)
		// 再修改权限
		os.Chmod(folderPath, 0777)
	}
	return folderPath, folderName
}

func test_sheet(f *excelize.File, sheet SheetModel) {
	sheetName := sheet.CurSheetName
	//map关系获取
	map_titles := map[string]int{}
	for _, v := range sheet.Titles {
		map_titles[v.DataName] = v.Sequence
	}

	// Create a new sheet.
	index := f.NewSheet(sheetName)

	for _, v := range sheet.Titles {
		sheetPosition := Div(v.Sequence) + "1"
		f.SetCellValue(sheetName, sheetPosition, v.TitleName)
	}

	//lineNo:=1
	for lineNum, v := range sheet.Datas {
		// Set value of a cell.
		//clumnNum := 0
		for kk, vv := range v {
			//clumnNum++
			fmt.Println(kk)
			col_name := Div(map_titles[kk])
			if col_name == "" {
				continue
			}
			sheetPosition := Div(map_titles[kk]) + strconv.Itoa(lineNum+2)
			switch vv.(type) {
			case string:
				f.SetCellValue(sheetName, sheetPosition, vv.(string))
				break
			case int:
				f.SetCellValue(sheetName, sheetPosition, vv.(int))
				break
			case float64:
				f.SetCellValue(sheetName, sheetPosition, vv.(float64))
				break
			}
		}
	}
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
}

// Div 数字转字母
func Div(Num int) string {
	var (
		Str  string = ""
		k    int
		temp []int //保存转化后每一位数据的值，然后通过索引的方式匹配A-Z
	)
	//用来匹配的字符A-Z
	Slice := []string{"", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
		"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	if Num > 26 { //数据大于26需要进行拆分
		for {
			k = Num % 26 //从个位开始拆分，如果求余为0，说明末尾为26，也就是Z，如果是转化为26进制数，则末尾是可以为0的，这里必须为A-Z中的一个
			if k == 0 {
				temp = append(temp, 26)
				k = 26
			} else {
				temp = append(temp, k)
			}
			Num = (Num - k) / 26 //减去Num最后一位数的值，因为已经记录在temp中
			if Num <= 26 {       //小于等于26直接进行匹配，不需要进行数据拆分
				temp = append(temp, Num)
				break
			}
		}
	} else {
		return Slice[Num]
	}
	for _, value := range temp {
		Str = Slice[value] + Str //因为数据切分后存储顺序是反的，所以Str要放在后面
	}
	return Str
}

type student struct {
	id    int
	name  string
	score int
}

func test6(arrlist []interface{}) {
	//结构体数组存储多为学员信息


	for k,v:=range arrlist{
		fmt.Println(k,v)
	}

}


