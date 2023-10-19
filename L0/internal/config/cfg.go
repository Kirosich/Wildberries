package config

import (
	"Wildber/internal/database"
	"Wildber/internal/nats"
	"Wildber/internal/transport"
	"fmt"
)

type Config struct {
	Router *transport.Router
	DB     *database.Database
	Cache  *transport.Cache
}

func InitCfg() (Config, error) {
	var cfg Config
	var err error
	cfg.DB = database.Init()
	cfg.Cache, err = transport.LoadCache(cfg.DB)
	cfg.Router = transport.InitRouter(cfg.Cache)

	if err != nil {
		fmt.Println("Error while Loading Cache")
		return cfg, err
	}
	go nats.StartListening(cfg.DB)
	return cfg, err
}
