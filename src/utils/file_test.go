package utils

import (
	"fmt"
	"os"
	"testing"
)

func TestWriteDataToExcel(t *testing.T) {
	// filepath, sheetName string, data []string
	currentDir, _ := os.Getwd()
	var tests = []struct {
		filepath  string
		sheetName string
		data      []string
	}{
		{fmt.Sprintf("%s/test3.xlsx", currentDir), "Table 1", []string{"one#two#three", "one#two#three"}},
		{fmt.Sprintf("%s/test4.xlsx", currentDir), "Table 2", []string{"first#second", "third#fourth"}},
	}
	for _, test := range tests {
		if err := WriteDataToExcel(test.filepath, test.sheetName, test.data); err != nil {
			t.Errorf("WriteDataToExcel(%s) failed: %v", test.filepath, err)
		}

	}
}
