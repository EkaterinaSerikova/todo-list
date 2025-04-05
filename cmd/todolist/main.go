package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Task struct {
	ID          string `json:"id" validate:"required" check:"unique"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Status      string `json:"status" validate:"required,oneof=Новая 'В процессе' Завершена"`
}

type User struct {
	ID       string `json:"id" check:"unique"`
	Name     string `json:"name"  validate:"required"`
	Email    string `json:"email"  validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

var tasks []Task
var users []User

func GetTasks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, tasks)
}

func CreateTask(ctx *gin.Context) {
	var task Task
	if err := ctx.ShouldBindBodyWithJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	task.ID = uuid.New().String()

	tasks = append(tasks, task)
	ctx.JSON(http.StatusCreated, task)
}

func GetTaskByID(ctx *gin.Context) {
	id := ctx.Param("id")
	for _, task := range tasks {
		if task.ID == id {
			ctx.JSON(http.StatusOK, task)
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedTask Task
	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Title = updatedTask.Title
			tasks[i].Description = updatedTask.Description
			tasks[i].Status = updatedTask.Status
			ctx.JSON(http.StatusOK, tasks[i])
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func GetUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, users)
}

func GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	for _, user := range users {
		if user.ID == id {
			ctx.JSON(http.StatusOK, user)
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func CreateUser(ctx *gin.Context) {
	var user User

	if err := ctx.ShouldBindBodyWithJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	user.ID = uuid.New().String()

	users = append(users, user)
	ctx.JSON(http.StatusCreated, user)
}

func UpdateUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedUser User
	if err := ctx.ShouldBindJSON(&updatedUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, user := range users {
		if user.ID == id {
			users[i].Name = updatedUser.Name
			users[i].Email = updatedUser.Email
			users[i].Password = updatedUser.Password
			ctx.JSON(http.StatusOK, users[i])
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			ctx.JSON(http.StatusOK, gin.H{"message": "User deleted"})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func main() {
	router := gin.Default()

	// для управления списком задач
	router.GET("/tasks", GetTasks)          // Получить список всех задач
	router.POST("/tasks", CreateTask)       // Создать новую задачу
	router.GET("/tasks/:id", GetTaskByID)   // Получить информацию о задаче по ее ID
	router.PUT("/tasks/:id", UpdateTask)    // Обновить информацию о задаче по ее ID
	router.DELETE("/tasks/:id", DeleteTask) // Удалить задачу по ее ID

	// для управления списком пользователей
	router.GET("/users", GetUsers)           // Получить список всех пользователей
	router.POST("/users", CreateUser)        // Создать нового пользователя
	router.GET("/users/:id", GetUserById)    // Получить информацию о пользователе по его ID
	router.PUT("/users/:id", UpdateUserById) // Обновить информацию о пользователе по его ID
	router.DELETE("/users/:id", DeleteUser)  // Удалить пользователя по его ID

	router.Run(":8080")
}
