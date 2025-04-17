package main

import (
	"context"

	"github.com/EkaterinaSerikova/todo-list/internal/app"
	"github.com/EkaterinaSerikova/todo-list/internal/config"
	"github.com/EkaterinaSerikova/todo-list/internal/repository/dbstorage"
	"github.com/EkaterinaSerikova/todo-list/internal/repository/memstorage"
	"github.com/EkaterinaSerikova/todo-list/internal/server"
	"github.com/EkaterinaSerikova/todo-list/internal/service"
	"github.com/EkaterinaSerikova/todo-list/pkg/logger"
)

// Главный исполняемый файл

func main() {
	// читаем настройки конфигурации
	cfg, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}

	// настраиваем логгер
	log := logger.Get(cfg.Debug)
	log.Debug().Msg("logger was initialized")
	log.Debug().Str("host", cfg.Host).Int("port", cfg.Port).Send()

	// выбираем хранилище данных
	var repo service.Repository
	repo, err = dbstorage.New(context.Background(), cfg.DbDsn)
	if err != nil {
		log.Warn().Err(err).Msg("failed to connect to db, using in-memory storage instead")
		repo = memstorage.New()
	} else {
		if err := dbstorage.Migrations(cfg.DbDsn, cfg.MigratePath); err != nil {
			log.Warn().Err(err).Msg("failed to migrations, using in-memory storage instead")
			repo = memstorage.New()
		}
	}

	// создаем сервисы бизнес-логики
	userService := service.NewUserService(repo)
	taskService := service.NewTaskService(repo)
	server := server.New(*cfg, userService, taskService)

	// запускаем сервер и приложение
	app := app.NewApp(*cfg, server, repo)
	if err := app.StartApp(); err != nil {
		panic(err)
	}
}
