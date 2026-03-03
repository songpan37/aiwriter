package main

import (
	"log"

	"aiwriter/api/v1"
	"aiwriter/internal/config"
	"aiwriter/internal/handler"
	"aiwriter/internal/middleware"
	"aiwriter/internal/service"
	"aiwriter/internal/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	var store storage.Storage
	if cfg.UseLocalStorage {
		store = storage.NewLocalStorage("./data")
		log.Println("Using local storage")
	} else {
		s3Client, err := storage.NewS3Client(
			cfg.S3Endpoint,
			cfg.S3Region,
			cfg.S3Bucket,
			cfg.S3AccessKey,
			cfg.S3SecretKey,
		)
		if err != nil {
			log.Fatalf("Failed to initialize S3 client: %v", err)
		}
		store = s3Client
		log.Println("S3 client initialized")
	}

	services := service.NewServices(cfg, store)
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
