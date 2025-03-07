package middleware

import (
	"net/http"

	"github.com/assaabriiii/gin-cookie-auth/internal/repositories"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(userRepo *repositories.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the session cookie
		username, err := c.Cookie("session")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Check if the user exists
		_, exists := userRepo.FindByUsername(username)
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Attach the username to the context
		c.Set("username", username)
		c.Next()
	}
}
