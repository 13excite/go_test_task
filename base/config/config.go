package config

import (
	"fmt"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// DefultConfigPath is default path for app config from rpm package
const DefultConfigPath = "/etc/testsrv.yaml"

// Config structure described yaml structure from /etc/testsrv.yaml
type Config struct {
	Listen   string `yaml:"listen"`
	DbName   string `yaml:"dbname"`
	DbServer string `yaml:"dbserver"`
}

// GetConfig reading and parsing configuration yaml file
func (conf *Config) GetConfig(configPath string) {
	if configPath == "" {
		configPath = DefultConfigPath
	}
	yamlConfig, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(yamlConfig, &conf)
	if err != nil {
		fmt.Println("Unmarshaled err")
		log.Fatal(err)
	}
}
