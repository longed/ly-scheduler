package utils

import (
	"time"
)

func Weekday(date string) string {
	return time.Now().Weekday().String()
}

func GetYYYYMMDDDate() string {
	return time.Now().Format("2006-01-02")
}
