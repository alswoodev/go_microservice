package routes

import (
	"userfc/cmd/user/handler"
	"userfc/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userHandler handler.UserHandler, secret string) {
	router.Use(middleware.RequestLogger())

	// Public API
	router.GET("/ping", userHandler.Ping)
	router.POST("/v1/register", userHandler.Register)
	router.POST("/v1/login", userHandler.Login)

	// Private API
	router.Use(middleware.AuthMiddleware(secret))
	router.GET("/v1/me", userHandler.GetMyInfo)
}