package config

import (
	"flag"
	"os"
)

type Config struct {
	ServerAddr string
	DBPath     string
}

func ReadConfig() *Config {
	var cfg Config
	flag.StringVar(&cfg.ServerAddr, "a", "localhost:8080", "server address")
	flag.StringVar(&cfg.DBPath, "d", "gophkeeper.db", "path to sqlite db")
	flag.Parse()

	if sAddr := os.Getenv("SERVER_ADDR"); sAddr != "" {
		cfg.ServerAddr = sAddr
	}
	if dbPath := os.Getenv("DATA_BASE_PATH"); dbPath != "" {
		cfg.DBPath = dbPath
	}

	return &cfg
}
