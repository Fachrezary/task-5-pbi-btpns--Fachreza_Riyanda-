// middlewares/auth_middleware.go

package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"projectapi/helpers"
)

// AuthMiddleware is a middleware for JWT authentication.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		claims, err := helpers.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Attach user ID to context
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}

// GetUserIDFromContext retrieves the user ID from the Gin context.
func GetUserIDFromContext(c *gin.Context) uint {
	userID, exists := c.Get("user_id")
	if !exists {
		return 0
	}

	return userID.(uint)
}

// RequireAuthenticated is a middleware that checks if the user is authenticated.
func RequireAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := GetUserIDFromContext(c)
		if userID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			c.Abort()
			return
		}
		c.Next()
	}
}
