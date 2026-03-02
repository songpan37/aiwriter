package main

import (
	"log"

	"aiwriter/api/v1"
	"aiwriter/internal/config"
	"aiwriter/internal/handler"
	"aiwriter/internal/middleware"
	"aiwriter/internal/model"
	"aiwriter/internal/repository"
	"aiwriter/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	db, err := config.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	err = db.AutoMigrate(
		&model.User{},
		&model.WorkCategory{},
		&model.Work{},
		&model.Volume{},
		&model.Chapter{},
		&model.Scene{},
		&model.OptimizationStep{},
		&model.OptimizationRecord{},
		&model.PublishTask{},
		&model.Notification{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migration completed")

	repos := repository.NewRepositories(db)
	services := service.NewServices(repos, cfg)
	handlers := handler.NewHandlers(services)

	r := gin.Default()
	r.Use(middleware.Cors())
	r.Use(middleware.Recovery())

	apiV1 := r.Group("/api/v1")

	auth := apiV1.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}

	protected := apiV1.Group("")
	protected.Use(middleware.Auth(cfg.JWTSecret))
	{
		v1.RegisterProtectedRoutes(protected, handlers)
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	log.Printf("Server starting on port %s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
