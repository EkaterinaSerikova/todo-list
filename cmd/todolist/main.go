package main

import (
	"github.com/EkaterinaSerikova/todo-list/internal/app"
	"github.com/EkaterinaSerikova/todo-list/internal/config"
	"github.com/EkaterinaSerikova/todo-list/internal/repository/memstorage"
	"github.com/EkaterinaSerikova/todo-list/internal/server"
	"github.com/EkaterinaSerikova/todo-list/internal/service"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}

	repo := memstorage.New()

	userService := service.NewUserService(repo)
	taskService := service.NewTaskService(repo)
	server := server.New(*cfg, userService, taskService)

	app := app.NewApp(*cfg, server, repo)
	if err := app.StartApp(); err != nil {
		panic(err)
	}
}
