package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/laudryfadian/starter-golang/internal/handler"
	"github.com/laudryfadian/starter-golang/internal/middleware"
)

func SetupUserRoutes(r *gin.Engine, userHandler *handler.UserHandler, authMiddleware *middleware.AuthMiddleware) {
	auth := r.Group("/api/auth")
	{
		auth.POST("/register", userHandler.Register)
		auth.POST("/login", userHandler.Login)
	}

	api := r.Group("/api")
	api.Use(authMiddleware.AuthGuard())
	{
		api.GET("/profile", userHandler.GetProfile)
		api.PUT("/profile", userHandler.UpdateProfile)
		api.DELETE("/profile", userHandler.DeleteUser)
		api.GET("/users", userHandler.GetAllUsers)
	}
}
