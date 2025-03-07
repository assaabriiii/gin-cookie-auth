package main

import (
	"github.com/assaabriiii/gin-cookie-auth/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Register routes
	routes.RegisterAuthRoutes(router)

	// Start the server
	router.Run(":8080")
}
