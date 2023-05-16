package application

import (
	"github.com/darchlabs/backoffice/config"
	"github.com/darchlabs/backoffice/internal/storage"
)

type App struct {
	Config   *config.Config
	SqlStore storage.SQL
}

type Config struct {
	Config   *config.Config
	SqlStore storage.SQL
}

func New(conf *Config) (*App, error) {
	return &App{
		Config:   conf.Config,
		SqlStore: conf.SqlStore,
	}, nil
}
