package utils

import (
	"time"
)

var (
	WeekdayInChinese = [7]string{
		"星期日",
		"星期一",
		"星期二",
		"星期三",
		"星期四",
		"星期五",
		"星期六",
	}
)

func Weekday(date string) string {
	return time.Now().Weekday().String()
}

func GetYYYYMMDDDate() string {
	return time.Now().Format("2006-01-02")
}
