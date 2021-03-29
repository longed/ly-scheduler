package model

import (
	"fmt"
	"os"
	"testing"
)

func TestReadMemberRecordSlice(t *testing.T) {
	currentDir, _ := os.Getwd()
	var tests = []struct {
		filepath string
	}{
		{fmt.Sprintf("%s/test3.xlsx", currentDir)},
	}
	for _, test := range tests {
		mrSlice, err := ReadMemberRecordSlice(test.filepath)
		if err != nil {
			t.Errorf("ReadMemberRecordSlice(%s) failed: %v", test.filepath, err)
		}
		for _, mr := range mrSlice {
			fmt.Printf("%s\n", mr.toString())
		}
	}
}
