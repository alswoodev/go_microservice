package routes

import (
	"userfc/cmd/user/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userHandler handler.UserHandler) {
	// Public API
	router.GET("/ping", userHandler.Ping)
	router.POST("/v1/register", userHandler.Register)
	router.POST("/v1/login", userHandler.Login)

	// Private API
}