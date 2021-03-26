package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

var (
	configOptions ConfigOptions
)

type ConfigOptions struct {
	MainOptions Main `toml:"main"`
}

type Main struct {
	Desc           string `toml:"desc"`
	Year           string `toml:"year"`
	MemberFilePath string `toml:"memberFilePath"`
}

func readConfig(path string) (ConfigOptions, error) {
	configOptions := ConfigOptions{}
	if _, err := toml.DecodeFile(path, &configOptions); err != nil {
		return ConfigOptions{}, err
	}
	return configOptions, nil
}

func GetConfigOptions() ConfigOptions {
	// read config from file
	configFilePaths := []string{"./config.toml", "./src/config.toml", "./conf/config.toml", "../conf/config.toml"}
	found := false
	for _, v := range configFilePaths {
		if config, err := readConfig(v); err == nil {
			configOptions = config
			found = true
			break
		}
	}
	if !found {
		log.Fatal("cannot find config.toml file. exit program.")
	}
	return configOptions
}
