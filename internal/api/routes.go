package api

import (
    "github.com/gin-gonic/gin"
    "otp-login-service/internal/api/handlers"
    "otp-login-service/internal/storage"
)

func RegisterRoutes(r *gin.Engine, userStore *storage.UserStorage, rateLimiter *storage.RateLimiter, jwtSecret []byte) {
    // ست کردن وابستگی‌ها
    handlers.UserStore = userStore
    handlers.JwtSecret = jwtSecret

    // گروه auth
    auth := r.Group("/auth")
    {
        auth.POST("/request-otp", handlers.RequestOTPHandler)
        auth.POST("/verify-otp", handlers.VerifyOTPHandler)
    }

    // مسیرهای کاربران
    r.GET("/users", handlers.ListUsersHandler)
    r.GET("/users/:phone", handlers.GetUserByPhoneHandler)
    //r.GET("/test", handlers.RequestOTPHandler)

}



