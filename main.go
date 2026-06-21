package main

import (
	"fmt"
	"userfc/cmd/user/handler"
	"userfc/cmd/user/repository"
	"userfc/cmd/user/resource"
	"userfc/cmd/user/service"
	"userfc/cmd/user/usecase"
	"userfc/config"
	"userfc/infrastructure/log"
	"userfc/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	redis := resource.InitRedis(&cfg)
	db := resource.InitDB(&cfg)
	log.SetupLogger()

	userRepository := repository.NewUserRepository(redis, db)
	userService := service.NewUserService(userRepository)
	userUsecase := usecase.NewUserUsecase(userService)
	userHandler := handler.NewUserHandler(userUsecase)

	port := cfg.App.Port
	router := gin.Default()

	routes.SetupRoutes(router, *userHandler)

	router.Run(fmt.Sprintf(":%s", port)) // :{port} -> :8080
	log.Logger.Printf("Server running on port: %s", port)
}