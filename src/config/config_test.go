package config

import (
	"fmt"
	"testing"
)

func TestFirstDateOfMonth(t *testing.T) {
	var tests = []struct {
		date string
	}{
		{"2020-03-21"},
	}
	for _, test := range tests {
		fmt.Printf("%v", dayTimeByString(test.date))
	}
}

func TestLastDateOfMonth(t *testing.T) {
	var tests = []struct {
		date string
	}{
		{"2020-03-21"},
	}
	for _, test := range tests {
		fmt.Printf("%v", dayTimeAfterMonth(test.date))
	}
}

func TestOther(t *testing.T) {
	var tests = []struct {
		date string
	}{
		{"2020-03-21"},
		{"2004-02-21"},
		{"2021-02-21"},
	}
	for _, test := range tests {
		d := dayTimeAfterMonth(test.date).Sub(dayTimeByString(test.date)).Hours()
		fmt.Printf("d=%f\n", d)
	}

}

func TestDaysOfMonth(t *testing.T) {
	var tests = []struct {
		date      string
		dayNumber int
	}{
		{"2020-03-21", 31},
		{"2020-02-21", 29},
		{"2021-02-21", 28},
	}

	for _, test := range tests {
		days := dayTimesOfMonth(test.date)
		if len(days) != test.dayNumber {
			t.Errorf("dayTimesOfMonth(%s) failed: %v", test.date, days)
		}
		fmt.Printf("")
	}
}

func TestDaysOfMemberSize(t *testing.T) {
	var tests = []struct {
		date       string
		memberSize int
	}{
		{"2018-03-21", 365},
		{"2019-02-21", 365},
		{"2020-02-21", 366},
		{"2021-02-21", 365},
		{"2022-02-21", 365},
		{"2004-02-21", 365},
	}
	for _, test := range tests {
		dayTimes := dayTimesOfMemberSize(test.date, test.memberSize)
		fmt.Printf("%s %d\n", test.date, len(dayTimes))
	}

}
