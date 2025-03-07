package controllers

import (
	"net/http"

	"github.com/assaabriiii/gin-cookie-auth/internal/models"
	"github.com/assaabriiii/gin-cookie-auth/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (ctrl *AuthController) Register(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := ctrl.authService.Register(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (ctrl *AuthController) Login(c *gin.Context) {
	var loginUser models.User
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user, err := ctrl.authService.Login(loginUser.Username, loginUser.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Set a cookie with the username
	c.SetCookie("session", user.Username, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged in successfully"})
}

func (ctrl *AuthController) Logout(c *gin.Context) {
	// Clear the session cookie
	c.SetCookie("session", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
