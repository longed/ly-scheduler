package model

import (
	"fmt"
	"ly-scheduler/src/config"
	"ly-scheduler/src/utils"
	"sort"
	"strconv"
	"strings"
	"time"
)

type MemberRecord struct {
	MemberName             string
	SchedulingStatus       int64 // 0 = schedule, 1 = not schedule, 2 = other
	Reason                 string
	SpecificWorkingWeekday string
}

type ScheduleRecord struct {
	Name         string
	TableContent map[string]string
	Ext          string // extend info
}

var (
	MrSlice []MemberRecord // data from member file
)

func init() {
	mfp := config.AllOptions.MainOptions.MemberFilePath
	mrSliceFromFile, err := ReadMemberRecordSlice(mfp)
	if err != nil {
		fmt.Printf("read member file error. filepath=%s, err=%v\n", mfp, err)
		return
	}
	MrSlice = append(MrSlice, mrSliceFromFile...)
}

func (sr *ScheduleRecord) String(beg, end int) string {
	kSlice := make([]string, 0, len(sr.TableContent))
	for k := range sr.TableContent {
		kSlice = append(kSlice, k)
	}
	sort.Strings(kSlice)

	vSlice := make([]string, 0, end-beg)
	for _, k := range kSlice[beg:end] {
		vSlice = append(vSlice, sr.TableContent[k])
	}
	return JoinStrings(utils.StringSeparator, sr.Name, JoinStrings(utils.StringSeparator, vSlice...), sr.Ext)
}

func (sr *ScheduleRecord) getTitle(beg, end int) string {
	kSlice := make([]string, 0, len(sr.TableContent))
	for k := range sr.TableContent {
		kTime, _ := time.Parse("2006-01-02", k)
		kSlice = append(kSlice, fmt.Sprintf("%s\n%s", k, utils.WeekdayInChinese[int(kTime.Weekday())]))
	}
	sort.Strings(kSlice)
	return JoinStrings(utils.StringSeparator, "姓名", JoinStrings(utils.StringSeparator, kSlice[beg:end]...), "备注")
}

func (mr *MemberRecord) toString() string {
	return JoinStrings(utils.CommaStringSeparator, mr.MemberName, strconv.FormatInt(mr.SchedulingStatus, 10), mr.Reason)
}

func (mr *MemberRecord) setDefaultValue() {
	mr.Reason = ""
	mr.SchedulingStatus = 0
	mr.MemberName = ""
}

func JoinStrings(sep string, elem ...string) string {
	return strings.Join(elem, sep)
}

func toTableContent(records []ScheduleRecord, beg, end int) []string {
	var data []string
	data = append(data, records[0].getTitle(beg, end))
	for _, record := range records {
		data = append(data, record.String(beg, end))
	}
	return data
}

// Save record to excel file
func SaveScheduleResult(filepath string, records []ScheduleRecord) error {
	pageSize := 7
	size := 0
	if len(records) > 0 {
		size = len(records[0].TableContent)
	}

	var data []string
	loop := size / pageSize
	beg := 0
	end := 0
	for i := 0; i < loop; i++ {
		beg = i * pageSize
		end = (i + 1) * pageSize
		data = append(data, toTableContent(records, beg, end)...)
		data = append(data, "") // blank line
	}

	if end < size {
		data = append(data, toTableContent(records, end, size)...)
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
	// skip sheet first line because first line is title.
	for _, row := range rows[1:] {
		mr := MemberRecord{}
		mr.setDefaultValue()
		if len(row) >= 1 {
			mr.MemberName = row[0]
		}
		if len(row) >= 2 {
			mr.SchedulingStatus, _ = strconv.ParseInt(row[1], 10, 32)
		}
		if len(row) >= 3 {
			mr.Reason = row[2]
		}
		if len(row) >= 4 {
			mr.SpecificWorkingWeekday = JoinStrings(utils.CommaStringSeparator, row[3:]...)
		}

		mrSlice = append(mrSlice, mr)
	}
	return mrSlice, nil
}
