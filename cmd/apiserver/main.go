package main

import (
	"flag"

	"github.com/avg73/realworld-go-api/internal/apiserver"
	"github.com/avg73/realworld-go-api/internal/config"
)

var (
	configPath string
)

func main() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
	flag.Parse()
	cfg := config.New(configPath)
	api := apiserver.New(cfg)
	api.Start()
}
