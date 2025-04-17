package app

import (
	"github.com/EkaterinaSerikova/todo-list/internal/config"
	"github.com/EkaterinaSerikova/todo-list/internal/server"
	"github.com/EkaterinaSerikova/todo-list/internal/service"
)

// определяем структуру и базовую логику управления приложением

type App struct {
	cfg       config.Config
	ServerApi *server.ServerApi
}

// создаем конструктор приложения
func NewApp(cfg config.Config, server *server.ServerApi, repo service.Repository) *App {
	return &App{
		cfg:       cfg,
		ServerApi: server,
	}
}

// реализуем точку входа для запуска приложения
func (app *App) StartApp() error {
	if err := app.ServerApi.Start(); err != nil {
		return err
	}
	return nil
}
