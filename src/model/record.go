package model

import (
	"ly-scheduler/src/utils"
	"strings"
)

type MemberRecord struct {
	MemberName   string
	IsScheduling bool
	Reason       string
}

type ScheduleRecord struct {
	Name  string
	Date1 string
	Date2 string
	Date3 string
	Date4 string
	Date5 string
	Date6 string
	Date7 string
	Ext   string // extend info
}

func (sr *ScheduleRecord) toString() string {
	return JoinStrings(utils.StringSeparator, sr.Name, sr.Date1, sr.Date2, sr.Date3, sr.Date4, sr.Date5, sr.Date6,
		sr.Date7, sr.Ext)
}

func JoinStrings(sep string, elem ...string) string {
	return strings.Join(elem, sep)
}

// Save record to excel file
func SaveScheduleResult(filepath string, records []ScheduleRecord) error {
	var data []string
	for _, record := range records {
		data = append(data, record.toString())
	}
	return utils.WriteDataToExcel(filepath, "Sheet1", data)
}
