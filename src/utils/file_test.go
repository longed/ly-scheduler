package utils

import (
	"testing"
)

func TestWriteDataToExcel(t *testing.T) {
	// filepath, sheetName string, data []string
	var tests = []struct {
		filepath  string
		sheetName string
		data      []string
	}{
		{"/home/lu/work/personal/go/ly-scheduler/src/test1.xlsx", "Table 1", []string{"今天天气真好呀！#可惜明天要去上班了。,学好 Golang 武装自己!", "今天天气真好呀！,可惜明天要去上班了。,学好 Golang 武装自己!"}},
		{"/home/lu/work/personal/go/ly-scheduler/src/test2.xlsx", "Table 2", []string{"好好学习#天天向上", "好好学习,天天向上"}},
	}
	for _, test := range tests {
		if err := WriteDataToExcel(test.filepath, test.sheetName, test.data); err != nil {
			t.Errorf("WriteDataToExcel(%s) failed: %v", test.filepath, err)
		}
	}
}
