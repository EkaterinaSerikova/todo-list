package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/EkaterinaSerikova/todo-list/internal/config"
	"github.com/EkaterinaSerikova/todo-list/pkg/logger"
)

type ServerApi struct {
	server *http.Server        // HTTP-сервер
	valid  *validator.Validate // Валидатор запросов
	repo   any                 // Репозиторий данных (интерфейсный тип)
}

func New(cfg config.Config, repo any) *ServerApi {
	// Инициализация HTTP-сервера с адресом из конфига
	server := http.Server{
		Addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), // Форматирование адреса типа "host:port"
	}

	// Создание роутера Gin с middleware по умолчанию (логгер и recovery)
	router := gin.Default()

	// Регистрация обработчиков для задач:

	// GET /tasks - получение списка задач
	router.GET("/tasks", func(c *gin.Context) {})
	// POST /tasks - создание новой задачи
	router.POST("/tasks", func(c *gin.Context) {})

	// Группа маршрутов для операций с конкретной задачей (/tasks/:id)
	task := router.Group("/tasks")
	{
		// PUT /tasks/:id - обновление задачи
		task.PUT("/:id", func(c *gin.Context) {})
		// DELETE /tasks/:id - удаление задачи
		task.DELETE("/:id", func(c *gin.Context) {})
		// GET /tasks/:id - получение задачи по ID
		task.GET("/:id", func(c *gin.Context) {})
	}

	// Группа маршрутов для пользователей (/user)
	user := router.Group("/user")
	{
		// POST /user/login - аутентификация пользователя
		user.POST("/login", func(c *gin.Context) {})
		// POST /user/register - регистрация нового пользователя
		user.POST("/register", func(c *gin.Context) {})
		// GET /user/profile - получение профиля пользователя
		user.GET("/profile", func(c *gin.Context) {})
	}

	// Назначение роутера обработчиком HTTP-сервера
	server.Handler = router

	// Возврат нового экземпляра ServerApi
	return &ServerApi{
		server: &server,         // Указатель на HTTP-сервер
		valid:  validator.New(), // Инициализация валидатора
		repo:   repo,            // Переданный репозиторий
	}
}

func (s *ServerApi) Start() error {
	log := logger.Get()

	log.Info().Str("server adress", s.server.Addr).Msg("server was started")

	return s.server.ListenAndServe()
}
