package main

import (
	"github.com/assaabriiii/gin-cookie-auth/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.RegisterAuthRoutes(router)
	router.Run(":8080")
}
