package model

import (
	"ly-scheduler/src/utils"
	"strconv"
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

func (mr *MemberRecord) toString() string {
	return JoinStrings(utils.CommaStringSeparator, mr.MemberName, strconv.FormatBool(mr.IsScheduling), mr.Reason)
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

// Read MemberRecordSlice from excel
func ReadMemberRecordSlice(filepath string) ([]MemberRecord, error) {
	rows, err := utils.ReadDataFromExcelSheet1(filepath)
	if err != nil {
		return []MemberRecord{}, err
	}

	var mrSlice []MemberRecord
	for _, row := range rows {
		mr := MemberRecord{}
		if len(row) >= 1 {
			mr.MemberName = row[0]
		}
		if len(row) >= 2 {
			mr.IsScheduling, _ = strconv.ParseBool(row[1])
		}
		if len(row) >= 3 {
			mr.Reason = row[2]
		}

		mrSlice = append(mrSlice, mr)
	}
	return mrSlice, nil
}
