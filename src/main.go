package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"ly-scheduler/src/config"
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

	// read config from file
	configFilePaths := []string{"./config.toml"}
	found := false
	for _, v := range configFilePaths {
		if config, err := config.ReadConfig(v); err == nil {
			configOptions = config
			found = true
		}
	}

	if !found {
		log.Fatal("cannot find config.toml file. exit program.")
	}
}

func main() {

}
