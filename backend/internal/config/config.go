package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort      string
	JWTSecret       string
	AIAPIKey        string
	S3Endpoint      string
	S3Region        string
	S3Bucket        string
	S3AccessKey     string
	S3SecretKey     string
	UseLocalStorage bool
}

func Load() *Config {
	godotenv.Load()

	return &Config{
		ServerPort:      getEnv("SERVER_PORT", "8080"),
		JWTSecret:       getEnv("JWT_SECRET", "your-secret-key"),
		AIAPIKey:        getEnv("AI_API_KEY", ""),
		S3Endpoint:      getEnv("S3_ENDPOINT", "obs.cn-north-4.myhuaweicloud.com"),
		S3Region:        getEnv("S3_REGION", "cn-north-4"),
		S3Bucket:        getEnv("S3_BUCKET", "aiwriter-content"),
		S3AccessKey:     getEnv("S3_ACCESS_KEY", "HPUA6DXH1UEYENA0NGEO"),
		S3SecretKey:     getEnv("S3_SECRET_KEY", "HTx7WLVnpbyR6PQaDhVNw6UhOoqEN5kkWJvdLBxB"),
		UseLocalStorage: getEnv("USE_LOCAL_STORAGE", "true") == "true",
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
