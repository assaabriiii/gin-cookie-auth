package routes

import (
	"net/http"

	"github.com/assaabriiii/gin-cookie-auth/internal/controllers"
	"github.com/assaabriiii/gin-cookie-auth/internal/middleware"
	"github.com/assaabriiii/gin-cookie-auth/internal/repositories"
	"github.com/assaabriiii/gin-cookie-auth/internal/services"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine) {
	// Initialize dependencies
	userRepo := repositories.NewUserRepository()
	authService := services.NewAuthService(userRepo)
	authController := controllers.NewAuthController(authService)

	// Public routes
	router.POST("/register", authController.Register)
	router.POST("/login", authController.Login)

	// Protected routes
	authGroup := router.Group("/")
	authGroup.Use(middleware.AuthMiddleware(userRepo))
	{
		authGroup.GET("/protected", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "You are authenticated!"})
		})
		authGroup.POST("/logout", authController.Logout)
	}
}
