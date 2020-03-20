package src

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

// GetLngAndLat 使用excelize处理excel.
func GetLngAndLat(path string){
	// 读取excel
	f ,err := excelize.OpenFile(path)
	if err != nil{
		fmt.Println("文件打开失败~, ",err)
	}

	// 获取当前sheet名
	name := f.GetSheetName(1)

	// 获取单元格的地址
	rows,err := f.Rows(name)
	if err != nil {
		fmt.Println("获取rows失败",err)
	}

	// 设置起始行
	beginRow := 1
	for rows.Next() {
		row, err := rows.Columns()
		if err != nil {
			println(err.Error())
		}
		// 拿出第二列的值
		addinfo ,err := GetAddressJSON(row[1])
		if err != nil {
			fmt.Println(err)
		}
		// 第三列，经度
		lng := fmt.Sprintf("C%d",beginRow)
		f.SetCellValue(name,lng,addinfo.Result.Location.Lng)

		// 第四轮，维度
		lat := fmt.Sprintf("D%d",beginRow)
		f.SetCellValue(name,lat,addinfo.Result.Location.Lat)

		f.Save()
		beginRow++
	}
}
