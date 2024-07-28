package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go_test/internal/handler"
	"go_test/internal/repository"
	"go_test/internal/service"
	"log"
)

// @title Clean Architecture Example API
// @version 1.0
// @description This is a sample server for a clean architecture example.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	r := gin.Default()

	repo := repository.NewRepository()
	svc := service.NewService(repo)
	h := handler.NewHandler(svc)

	v1 := r.Group("/api/v1")
	{
		v1.GET("/entity", h.GetMethod)
		v1.GET("/entity/:id", h.GetMethodById)
		v1.POST("/entity", h.PostMethod)
		v1.PUT("/entity/:id", h.PutMethod)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	fmt.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
