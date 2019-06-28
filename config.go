package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port int
	}
	Jwt struct {
		Key            string
		ExpirationTime int
	}
	Credentials map[string]string
}

func loadConfig(location string) (*Config, error) {

	cfg := &Config{}
	_, err := os.Stat(location)

	if os.IsNotExist(err) {
		return cfg, fmt.Errorf("Config file not found: %s", location)
	}

	yamlFile, err := ioutil.ReadFile(location)
	if err != nil {
		return cfg, fmt.Errorf("Load config error %v ", err)
	}

	err = yaml.Unmarshal(yamlFile, cfg)
	if err != nil {
		return cfg, fmt.Errorf("Unmarshal: %v", err)
	}

	if cfg.Server.Port == 0 {
		return cfg, fmt.Errorf("No cfg.Server.Port value")

	}
	if cfg.Jwt.ExpirationTime == 0 {
		return cfg, fmt.Errorf("No cfg.Jwt.ExpirationTime value")
	}
	if cfg.Jwt.Key == "" {
		return cfg, fmt.Errorf("No cfg.Jwt.Key value")
	}
	if len(cfg.Credentials) < 1 {
		return cfg, fmt.Errorf("No cfg.Credentials")
	}

	return cfg, err
}
