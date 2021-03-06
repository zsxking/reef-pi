package main

import (
	"github.com/ranjib/reef-pi/api"
	"github.com/ranjib/reef-pi/controller"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Controller controller.Config `yaml:"controller"`
	API        api.ServerConfig  `yaml:"api"`
}

var DefaultConfig = Config{
	Controller: controller.DefaultConfig,
	API:        api.DefaultConfig,
}

func ParseConfig(filename string) (*Config, error) {
	var c Config
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(content, &c); err != nil {
		return nil, err
	}
	for k, o := range c.Controller.Outlets {
		o.Name = k
		if o.Type == "" {
			o.Type = "switch"
		}
		c.Controller.Outlets[o.Name] = o
	}
	return &c, nil
}
