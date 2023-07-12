package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
}

func New(configPath string) *Config {
	cfg := new(Config)
	_, err := toml.DecodeFile(configPath, cfg)
	if err != nil {
		log.Fatal("toml decode error: ", err)
	}
	return cfg
}
