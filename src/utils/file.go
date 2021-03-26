package utils

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func WriteToExcel() {
	excelFile := excelize.NewFile()
	index := excelFile.NewSheet("Sheet1")
	excelFile.NewSheet("Sheet2")
	excelFile.SetCellValue("Sheet2", "A2", "Hello world.")
	excelFile.SetCellValue("Sheet1", "B2", 100)
	excelFile.SetActiveSheet(index)

	if err := excelFile.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}

}
