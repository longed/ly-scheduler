package utils

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

const (
	StringSeparator      = `#`
	CommaStringSeparator = `,`
)

// Write data to excel file.
// If filepath already exist file, delete old file create new one.
// Line of data can contains multiple columns, join columns by StringSeparator.
func WriteDataToExcel(filepath, sheetName string, data []string) error {
	var f *excelize.File
	var err error

	// file exist then delete it.
	if _, err = os.Stat(filepath); !errors.Is(err, os.ErrNotExist) {
		os.Remove(filepath)
	}

	f = excelize.NewFile()
	defer f.SaveAs(filepath)

	f.NewSheet(sheetName)

	for lineNumber, line := range data {
		// split every line to get columns
		cols := strings.Split(line, StringSeparator)

		for colNumber, col := range cols {
			// row and column index start from 1
			cellName, err := excelize.CoordinatesToCellName(colNumber+1, lineNumber+1)
			if err != nil {
				fmt.Printf("convert coordinate to cell name error, err=%v, can't write value to file\n", err)
				continue
			}
			f.SetCellValue(sheetName, cellName, col)
		}
	}
	return nil
}

// Read data from excel file, put all data in [][]string data type.
func ReadDataFromExcel(filepath, sheetName string) ([][]string, error) {
	if len(strings.TrimSpace(filepath)) == 0 {
		return [][]string{}, fmt.Errorf("filepath is blank")
	}

	if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
		return [][]string{}, fmt.Errorf("file not exist. filepath=%s", filepath)
	}

	var f *excelize.File
	var err error
	if f, err = excelize.OpenFile(filepath); err != nil {
		return [][]string{}, fmt.Errorf("open file errpr, err=%v", err)
	}
	return f.GetRows(sheetName)
}

// Read data from sheet "Sheet1"
func ReadDataFromExcelSheet1(filepath string) ([][]string, error) {
	return ReadDataFromExcel(filepath, "Sheet1")
}
