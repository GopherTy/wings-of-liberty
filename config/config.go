package config

import (
	"encoding/json"
	"log"
	"os"

	"go.uber.org/zap"
)

// config gobal configuration object
var config Config

// Config configuration object
type Config struct {
	Freedom Freedom
	Logger  *zap.Logger
}

// Freedom cross greate firewall configuration object
type Freedom struct {
	Port int    `json:"Port"`
	Addr string `json:"Addr"`
}

//	initialize configuration object
func init() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("initializate logger fail %v", err)
	}
	file, err := os.Open(CONFIGPATH)
	if err != nil {
		log.Fatalf("initializate config fail %v", err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("decode freedom.json file fail %v", err)
	}
	config.Logger = logger
}

// GetConfig get a configuration object
func GetConfig() (configuration *Config) {
	configuration = &config
	return
}
