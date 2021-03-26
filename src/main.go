package main

import (
	"fmt"
	"io/ioutil"
	"ly-scheduler/src/config"
	
)

func init() {
	bannerPaths := []string{"./banner.txt", "../doc/banner.txt", "./doc/banner.txt", "../misc/doc/banner.txt"}
	for _, v := range bannerPaths {
		data, err := ioutil.ReadFile(v)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(data))
			break
		}
	}

	configFilePaths := []string{"./config.toml"}
	for _, v := range configFilePaths {
		if configOptions, err:=config.ReadConfig(v); err == nil {
			fmt.Println(configOptions.MainOptions.Desc)
		}

	}
}

func main() {

}
