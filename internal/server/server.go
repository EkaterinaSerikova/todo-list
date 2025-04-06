package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/EkaterinaSerikova/todo-list/internal/config"
	"github.com/EkaterinaSerikova/todo-list/internal/service"
)

type ServerApi struct {
	server   *http.Server
	valid    *validator.Validate
	uService *service.UserService
	tService *service.TaskService
}

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
	s.configRoutes()
	return s.server.ListenAndServe()
}

func (s *ServerApi) configRoutes() {
	router := gin.Default()
	router.GET("/tasks", s.getTasks)
	router.POST("/tasks", s.createTask)

	task := router.Group("/tasks")
	{
		task.GET("/:id", s.getTaskById)
		task.PUT("/:id", s.updateUserById)
		task.DELETE("/:id", s.deleteTask)
	}

	router.GET("/users", s.getUsers)
	router.POST("/users", s.createUser)

	users := router.Group("/users")
	{
		users.POST("/register", s.registerUser)
		users.POST("/login", s.loginUser)

		users.GET("/users/:id", s.getUserById)
		users.PUT("/users/:id", s.updateUserById)
		users.DELETE("/users/:id", s.deleteUser)
	}

	s.server.Handler = router
}
