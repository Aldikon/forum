package config

import (
	"encoding/json"
	"os"
)

const CONFIG_FILE_NAME = "./config/config.json"

// write config
type Config struct {
	Server struct {
		Port string `json:"port"`
	} `json:"server"`
	Path struct {
		DB       string `json:"db"`
		Template string `json:"template"`
		Static   string `jsin:"static"`
	} `json:"path"`
}

// Global config.
var C Config

// Write config.
func ReadConfig() error {
	// t := http.Server{
	// 	WriteTimeout:   1,
	// 	ReadTimeout:    1,
	// 	MaxHeaderBytes: 1,
	// }
	data, err := os.ReadFile(CONFIG_FILE_NAME)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &C)
}
