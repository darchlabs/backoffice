package application

import "github.com/darchlabs/backoffice/internal/storage"

type App struct {
	SqlStore storage.SQL
}

type Config struct {
	SqlStore storage.SQL
}

func New(conf *Config) (*App, error) {
	return &App{
		SqlStore: conf.SqlStore,
	}, nil
}
