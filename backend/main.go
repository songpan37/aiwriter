package main

import (
	"log"

	"aiwriter/api/v1"
	"aiwrit
	"aiwriter/internal/config"
	"aiwriter/internal/handler"
	"aiwriter/internal/middleware"
	"aiwriter/internal/repository"
	"aiwriter/api/v1"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	db, err := config.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	repos := repository.NewRepositories(db)
	services := service.NewServices(repos, cfg)
	handlers := handler.NewHandlers(services)

	r := gin.Default()
	r.Use(middleware.Cors())
	r.Use(middleware.Recovery())

	apiV1 := r.Group("/api/v1")
	apiV1.Use(middleware.Auth(cfg.JWTSecret))
	{
		v1.RegisterRoutes(apiV1, handlers)
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	log.Printf("Server starting on port %s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
