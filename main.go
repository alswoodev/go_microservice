package main

import (
	"fmt"
	"userfc/cmd/user/handler"
	"userfc/config"
	"userfc/infrastructure/log"
	"userfc/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	log.SetupLogger()

	port := cfg.App.Port
	router := gin.Default()

	userHandler := handler.NewUserHandler()
	routes.SetupRoutes(router, *userHandler)

	router.Run(fmt.Sprintf(":%s", port)) // :{port} -> :8080
	log.Logger.Printf("Server running on port: %s", port)
}