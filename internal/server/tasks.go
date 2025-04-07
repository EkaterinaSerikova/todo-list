package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/EkaterinaSerikova/todo-list/internal/domain/errors"
	"github.com/EkaterinaSerikova/todo-list/internal/domain/models"
)

func (s *ServerApi) getTasks(c *gin.Context) {
	tasks, err := s.tService.GetTasks()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (s *ServerApi) createTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindBodyWithJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := s.tService.CreateTask(task); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (s *ServerApi) getTaskById(c *gin.Context) {
	uid := c.Param("id")
	tasks, err := s.tService.GetTasks()
	for _, task := range tasks {
		if task.UID == uid {
			c.JSON(http.StatusOK, task)
			return
		}
	}
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": errors.ErrTaskNotFound})
		return
	}
}

func (s *ServerApi) updateTask(c *gin.Context) {
	id := c.Param("id")
	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tasks, err := s.tService.GetTasks()

	for i, task := range tasks {
		if task.UID == id {
			tasks[i].Title = updatedTask.Title
			tasks[i].Description = updatedTask.Description
			tasks[i].Status = updatedTask.Status
			c.JSON(http.StatusOK, tasks[i])
			return
		}
	}
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": errors.ErrTaskNotFound})
		return
	}
}

func (s *ServerApi) deleteTask(c *gin.Context) {
	id := c.Param("id")

	tasks, err := s.tService.GetTasks()
	for i, task := range tasks {
		if task.UID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
			return
		}
	}
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": errors.ErrTaskNotFound})
		return
	}
}
