package pkg

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

var GlobalConfig = Config{}

func init() {
	yamlFile, err := os.ReadFile("/Users/Mr_J/Desktop/synchron-space/conf/conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = yaml.Unmarshal(yamlFile, &GlobalConfig)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

type Config struct {
	WatcherPath string `yaml:"watcher-path"`
}
