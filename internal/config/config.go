package config

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	Path   string
	Port   int64
	Debug  bool
	Routes map[string]RouteConfig
}

type RouteConfig struct {
	Status *int
	Body   map[string]interface{}
}

func Load(configPath string, expectedFileName string) (*AppConfig, error) {
	stats, err := os.Stat(configPath)
	if err != nil {
		return nil, err
	}

	if stats.IsDir() {
		configPath = filepath.Join(configPath, expectedFileName)
	}

	yamlConfig, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	log.Printf("Loading config file %s", configPath)

	parsedConfig := AppConfig{}

	err = yaml.Unmarshal(yamlConfig, &parsedConfig)
	if err != nil {
		return nil, err
	}

	parsedConfig.Path = configPath

	return &parsedConfig, nil
}
