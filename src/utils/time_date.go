package utils

import (
	"time"
)

func Weekday(date string) string {
	return time.Now().Weekday().String()
}
