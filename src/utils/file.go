package utils

import (
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"os"
	"strings"
)

const (
	StringSeparator = `#`
)

// write data to excel file
// if filepath already exist file, delete old file create new one.
// line of data can contains multiple columns, join columns by StringSeparator.
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
		// split every line to get colmuns
		cols := strings.Split(line, StringSeparator)

		for colNumber, col := range cols {
			// row and column index start from 1
			// first row cannot be used
			cellName, err := excelize.CoordinatesToCellName(colNumber+1, lineNumber+2)
			if err != nil {
				fmt.Printf("convert coordinate to cell name error, err=%v, can't write value to file\n", err)
				continue
			}
			f.SetCellValue(sheetName, cellName, col)
		}
	}
	return nil
}
