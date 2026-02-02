package handler

import (
	"IronOps/internal/model"
	"IronOps/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func LoginHandler(c *gin.Context) {
	// Simple mock login
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := service.GetUserByUsername(creds.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// In real world, check password hash. Here just return user info.
	c.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"user":    user,
	})
}

func ListUsersHandler(c *gin.Context) {
	users, err := service.ListUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func ListRolesHandler(c *gin.Context) {
	roles := []model.RoleType{model.RoleAdmin, model.RoleOps, model.RoleViewer}
	c.JSON(http.StatusOK, roles)
}
