package helper

import (
	"encoding/json"
	"io/ioutil"
)

// Config holds all configuration data
type Config struct {
	WeightPath    string `json:"weightdata"`
	DefaultHeight string `json:"defaultheight"`
}

// GetConfig to read config from json
func GetConfig(filename string) Config {
	config := Config{}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	return config
}
