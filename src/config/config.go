package config

import(
	"github.com/BurntSushi/toml"
)

type ConfigOptions struct {
	MainOptions MainOptions `toml:"main"`
}

type MainOptions struct {
	Desc string `toml:"desc"`
}

func ReadConfig(path string) (ConfigOptions ,error){
	configOptions := ConfigOptions{}
	if _, err := toml.DecodeFile(path, &configOptions); err != nil {
		return ConfigOptions{}, err
	}
	return configOptions, nil
}