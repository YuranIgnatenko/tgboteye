package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Token string `json:"Token"`
}

func New(nameFile string) *Config {
	conf := Config{}
	f, err := ioutil.ReadFile(nameFile)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal([]byte(f), &conf)

	return &conf
}
