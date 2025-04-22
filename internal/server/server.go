package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/EkaterinaSerikova/todo-list/internal/config"
	"github.com/EkaterinaSerikova/todo-list/internal/service"
)

// реализация HTTP-сервера на фреймворке Gin
// принимает HTTP-запросы и делегирует обработку сервисам (UserService, TaskService)

// создание экземпляра ServerApi с настройками из конфига (host, port)
type ServerApi struct {
	server   *http.Server
	valid    *validator.Validate
	uService *service.UserService
	tService *service.TaskService
}

// конструктор для создания экземпляра ServerApi
func New(cfg config.Config, uService *service.UserService, tService *service.TaskService) *ServerApi {
	server := http.Server{
		Addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
	}

	return &ServerApi{
		server:   &server,
		valid:    validator.New(),
		uService: uService,
		tService: tService,
	}
}

func (s *ServerApi) Start() error {
	s.configRoutes()                 // настройка роутов
	return s.server.ListenAndServe() // запуск сервера
}

// настройка маршрутов
func (s *ServerApi) configRoutes() {
	router := gin.Default()
	router.GET("/tasks", s.getTasks)
	router.POST("/tasks", s.createTask)

	task := router.Group("/tasks")
	{
		task.POST("/save-tasks", s.saveTasks)
		task.GET("/:id", s.getTaskById)
		task.PUT("/:id", s.updateTask)
		task.DELETE("/:id", s.deleteTask)
	}

	router.GET("/users", s.getUsers)

	users := router.Group("/users")
	{
		users.POST("/register", s.registerUser)
		users.POST("/login", s.loginUser)

		users.GET("/:id", s.getUserById)
		users.PUT("/:id", s.updateUserById)
		users.DELETE("/:id", s.deleteUser)
	}

	s.server.Handler = router
}
