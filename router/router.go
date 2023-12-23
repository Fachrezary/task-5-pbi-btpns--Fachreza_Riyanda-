package router

import (
	"github.com/gin-gonic/gin"
	"projectapi/controllers"
	"projectapi/middlewares"
)

// SetupRouter configures the Gin router and sets up the endpoints.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Group routes that require authentication
	authRoutes := r.Group("/api")
	authRoutes.Use(middlewares.AuthMiddleware())
	{
		authRoutes.GET("/user/:userId", controllers.GetUserByID)
		authRoutes.PUT("/user/:userId", controllers.UpdateUser)
		authRoutes.DELETE("/user/:userId", controllers.DeleteUser)

		authRoutes.POST("/photo", controllers.CreatePhoto)
		authRoutes.GET("/photos", controllers.GetPhotos)
		authRoutes.PUT("/photo/:photoId", controllers.UpdatePhoto)
		authRoutes.DELETE("/photo/:photoId", controllers.DeletePhoto)
	}

	// Public routes
	r.POST("/users/register", controllers.CreateUser)

	r.POST("/users/login", controllers.LoginUser)

	return r
}
