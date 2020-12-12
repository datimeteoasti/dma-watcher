package dmawatcher

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	Watcher WatcherConfig `yaml:"watcher"`
}

type WatcherConfig struct {
	Path            string `yaml:"path"`
	RegexFilterHook string `yaml:"regexFilterHook"`
}

func getConf() (*AppConfig, error) {
	var c *AppConfig
	yamlFile, err := ioutil.ReadFile("config/app.yml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
