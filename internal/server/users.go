package server

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	user_errors "github.com/EkaterinaSerikova/todo-list/internal/domain/errors"
	"github.com/EkaterinaSerikova/todo-list/internal/domain/models"
)

// HTTP-обработчики для работы с пользователями

func (s *ServerApi) registerUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid request body",
			"details": err.Error(), // Опционально, можно убрать в production
		})
		return
	}

	uid, err := s.uService.RegisterUser(user)
	if err != nil {
		switch {
		case errors.Is(err, user_errors.ErrUserAlreadyExists):
			c.JSON(http.StatusConflict, gin.H{
				"error": "user already exists",
			})
		case errors.Is(err, user_errors.ErrInvalidEmail):
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid email format",
			})
		case errors.Is(err, user_errors.ErrWeakPassword):
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "password does not meet requirements",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to register user",
			})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"uid":     uid,
		"message": "user registered successfully",
	})
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
		switch {
		case errors.Is(err, user_errors.ErrUserNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		case errors.Is(err, user_errors.ErrInvalidCredentials):
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		case errors.Is(err, user_errors.ErrAccountLocked):
			c.JSON(http.StatusForbidden, gin.H{"error": "account temporarily locked"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "authentication service unavailable"})
		}
		return
	}

	c.SetCookie("user_id", user_id, 3600, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"user_id": user_id})
}

func (s *ServerApi) getUsers(c *gin.Context) {
	users, err := s.uService.GetUsers()
	if err != nil {
		if errors.Is(err, user_errors.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusOK, []models.User{})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (s *ServerApi) getUserById(c *gin.Context) {
	uid := c.Param("id")
	users, err := s.uService.GetUsers()

	if err != nil {
		if errors.Is(err, user_errors.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}

	for _, user := range users {
		if user.UID == uid {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": user_errors.ErrUserNotFound})
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
	if err != nil {
		if errors.Is(err, user_errors.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}

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
		c.JSON(http.StatusNotFound, gin.H{"error": user_errors.ErrUserNotFound})
		return
	}
}

func (s *ServerApi) deleteUser(c *gin.Context) {
	id := c.Param("id")
	users, err := s.uService.GetUsers()
	if err != nil {
		if errors.Is(err, user_errors.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "users list not available"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not retrieve users"})
		}
		return
	}
	for i, user := range users {
		if user.UID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
			return
		}
	}
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": user_errors.ErrUserNotFound})
		return
	}
}
