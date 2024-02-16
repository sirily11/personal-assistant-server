package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"personal-assistant/internal/config/constants/environments"
)

// APIKeyMiddleware is responsible for validating the JWT token sent in the header of the request from
func APIKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("x-api-key")
		expectedApiKey := environments.AdminApiKey
		if apiKey != expectedApiKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		}
		c.Next()
	}
}
