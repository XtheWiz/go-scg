package config

import (
	"fmt"
	"go-scg/internal/model"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func LoadConfig(path string) (model.Config, error) {
	config := model.Config{}
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("yamlFile.Get err #%v ", err)
		return config, err
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Printf("error: %v", err)
		return config, err
	}
	return config, nil
}
