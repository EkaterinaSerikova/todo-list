package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/EkaterinaSerikova/todo-list/internal/domain/errors"
	"github.com/EkaterinaSerikova/todo-list/internal/domain/models"
)

// HTTP-обработчики для работы с пользователями

func (s *ServerApi) registerUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindBodyWithJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uid, err := s.uService.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"uid": uid})
}

func (s *ServerApi) loginUser(c *gin.Context) {
	var user models.UserRequest
	err := c.ShouldBindBodyWithJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user_id, err := s.uService.LoginUser(user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("user_id", user_id, 3600, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"user_id": user_id})
}

func (s *ServerApi) getUsers(c *gin.Context) {
	users, err := s.uService.GetUsers()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (s *ServerApi) getUserById(c *gin.Context) {
	uid := c.Param("id")
	users, err := s.uService.GetUsers()

	for _, user := range users {
		if user.UID == uid {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": errors.ErrUserNotFound})
		return
	}
}

func (s *ServerApi) updateUserById(c *gin.Context) {
	id := c.Param("id")
	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users, err := s.uService.GetUsers()

	for i, user := range users {
		if user.UID == id {
			users[i].Name = updatedUser.Name
			users[i].Login = updatedUser.Login
			users[i].Password = updatedUser.Password
			c.JSON(http.StatusOK, users[i])
			return
		}
	}
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": errors.ErrUserNotFound})
		return
	}
}

func (s *ServerApi) deleteUser(c *gin.Context) {
	id := c.Param("id")
	users, err := s.uService.GetUsers()
	for i, user := range users {
		if user.UID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
			return
		}
	}
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": errors.ErrUserNotFound})
		return
	}
}
