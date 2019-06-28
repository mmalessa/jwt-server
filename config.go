package main

import (
	"io/ioutil"
	"log"

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

func loadConfig(location string) *Config {

	yamlFile, err := ioutil.ReadFile(location)
	if err != nil {
		log.Printf("Load config ERROR  #%v ", err)
	}

	cfg := &Config{}
	err = yaml.Unmarshal(yamlFile, cfg)
	if err != nil {
		log.Printf("Unmarshal ERROR: %v", err)
	}

	return cfg
}
