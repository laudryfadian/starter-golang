//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/laudryfadian/starter-golang/internal/config"
	"github.com/laudryfadian/starter-golang/internal/handler"
	"github.com/laudryfadian/starter-golang/internal/middleware"
	"github.com/laudryfadian/starter-golang/internal/repository"
	"github.com/laudryfadian/starter-golang/internal/routes"
	"github.com/laudryfadian/starter-golang/internal/usecase"
)

func InitializeServer() (*gin.Engine, error) {
	wire.Build(
		config.ConnectDB,
		repository.NewUserRepository,
		usecase.NewUserUsecase,
		handler.NewUserHandler,
		middleware.NewAuthMiddleware,
		routes.SetupRoutes,
	)
	return nil, nil
}
