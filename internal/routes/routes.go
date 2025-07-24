package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/laudryfadian/starter-golang/internal/handler"
	"github.com/laudryfadian/starter-golang/internal/middleware"
)

func SetupRoutes(userHandler *handler.UserHandler, authMiddleware *middleware.AuthMiddleware) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	SetupUserRoutes(r, userHandler, authMiddleware)

	return r
}
