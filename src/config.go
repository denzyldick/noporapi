package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Config struct {
	MongoDBUrl string `json:"MONGODB_URL"`
	ServePort  string `json:"SERVE_PORT"`
}

// This file contains all credentials needed to
// connect to services
const ConfigFile = "./config.json"

func (p *Config) New() *Config {
	bytes, err := ioutil.ReadFile(ConfigFile)

	if err != nil {
		log.Fatal("- [X] There is no configuration file.")
	}
	e := json.Unmarshal(bytes, &p)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println("- [OK] Configuration file has been loaded. ")
	return p
}
