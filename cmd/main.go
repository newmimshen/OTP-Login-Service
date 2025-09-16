package main

import (
    "github.com/gin-gonic/gin"
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"
    "otp-login-service/docs" // ← برای بارگذاری کامنت‌ها
    "otp-login-service/internal/api"
    "otp-login-service/internal/storage"
    "github.com/joho/godotenv"
    "os"

)

// @title OTP Login API
// @version 1.0
// @description API for OTP-based login and user management
// @host localhost:8080
// @BasePath /

func main() {
    godotenv.Load()

    // مقداردهی به اطلاعات Swagger
    docs.SwaggerInfo.Title = "OTP Login API"
    docs.SwaggerInfo.Description = "API for OTP-based login and user management"
    docs.SwaggerInfo.Version = "1.0"
    docs.SwaggerInfo.Host = "localhost:8080"
    docs.SwaggerInfo.BasePath = "/"

    jwtSecret := os.Getenv("JWT_SECRET")
    if jwtSecret == "" {
        panic("JWT_SECRET not set in environment")
    }

    r := gin.Default()

    userStore := storage.NewUserStorage()
    rateLimiter := storage.NewRateLimiter()

    api.RegisterRoutes(r, userStore, rateLimiter, []byte(jwtSecret))

    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    r.Run(":8080")
}