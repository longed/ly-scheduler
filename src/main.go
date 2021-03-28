package main

import (
	"fmt"
	"io/ioutil"
	"ly-scheduler/src/config"
	"ly-scheduler/src/utils"
)

var (
	configOptions config.ConfigOptions
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

	configOptions = config.GetConfigOptions()
}

func main() {
	fmt.Println(utils.Weekday(""))
}
