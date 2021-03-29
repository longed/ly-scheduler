package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"ly-scheduler/src/utils"
)

var (
	configOptions ConfigOptions
)

type ConfigOptions struct {
	MainOptions Main `toml:"main"`
}

type Main struct {
	Desc           string `toml:"desc"`           // description of the program
	YearMonthDay   string `toml:"yearMonthDay"`   // scheduler start date
	MemberFilePath string `toml:"memberFilePath"` // file that contains members
}

// Set ConfigOptions fields default value.
func (configOptions *ConfigOptions) setDefaultValues() {
	configOptions.MainOptions.Desc = ""
	configOptions.MainOptions.YearMonthDay = utils.GetYYYYMMDDDate()
	configOptions.MainOptions.MemberFilePath = ""
}

func readConfig(path string) (ConfigOptions, error) {
	configOptions := ConfigOptions{}
	configOptions.setDefaultValues()
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
