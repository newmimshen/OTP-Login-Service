package handlers

// import (
//     "fmt"
//     "net/http"
//     "time"

//     "github.com/gin-gonic/gin"
//     "otp-login-service/internal/storage"
//     "otp-login-service/internal/utils"
// )

// // ✅ وابستگی‌ها به صورت global
// var RateLimiter *storage.RateLimiter

// type OTPRequest struct {
//     Phone string `json:"phone" binding:"required"`
// }

// // @Summary درخواست OTP برای شماره تلفن
// // @Tags Auth
// // @Accept json
// // @Produce json
// // @Param request body OTPRequest true "شماره تلفن"
// // @Success 200 {object} map[string]string
// // @Failure 400 {object} map[string]string
// // @Failure 429 {object} map[string]string
// // @Router /auth/request-otp [post]
// func RequestOTPHandler(c *gin.Context) {
//     var req OTPRequest

//     if err := c.ShouldBindJSON(&req); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": "شماره تلفن الزامی است"})
//         return
//     }

//     if !RateLimiter.Allow(req.Phone) {
//         c.JSON(http.StatusTooManyRequests, gin.H{"error": "محدودیت درخواست OTP: حداکثر ۳ بار در ۱۰ دقیقه"})
//         return
//     }

//     code := utils.GenerateOTP()
//     expiresAt := time.Now().Add(2 * time.Minute)
//     storage.SaveOTP(req.Phone, code, expiresAt)

//     fmt.Printf("OTP برای %s: %s\n", req.Phone, code)

//     c.JSON(http.StatusOK, gin.H{
//         "message": "OTP ارسال شد (در کنسول چاپ شده)",
//     })
// }
