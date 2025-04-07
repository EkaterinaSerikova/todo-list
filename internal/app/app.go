package app

import (
	"github.com/EkaterinaSerikova/todo-list/internal/config"
	"github.com/EkaterinaSerikova/todo-list/internal/server"
	"github.com/EkaterinaSerikova/todo-list/internal/service"
)

type App struct {
	cfg       config.Config
	ServerApi *server.ServerApi
	repo      service.Repository
}

func NewApp(cfg config.Config, server *server.ServerApi, repo service.Repository) *App {
	return &App{
		cfg:       cfg,
		ServerApi: server,
		repo:      repo,
	}
}

func (app *App) StartApp() error {
	if err := app.ServerApi.Start(); err != nil {
		return err
	}
	return nil
}
