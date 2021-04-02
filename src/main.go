package main

import (
	"fmt"
	"io/ioutil"
	"ly-scheduler/src/config"
	"ly-scheduler/src/model"
	"ly-scheduler/src/schedule"
)

func init() {
	// print banner
	bannerPaths := []string{"./banner.txt", "../doc/banner.txt", "./doc/banner.txt", "../misc/doc/banner.txt"}
	for _, v := range bannerPaths {
		data, err := ioutil.ReadFile(v)
		if err == nil {
			fmt.Println(string(data))
			break
		}
	}
}

func main() {
	srSlice := schedule.RsScheduler.DoSchedule("")
	err := model.SaveScheduleResult(config.AllOptions.MainOptions.ScheduleTableFilePath, srSlice)
	if err == nil {
		fmt.Printf("create schedule table success.\n")
	}
}
