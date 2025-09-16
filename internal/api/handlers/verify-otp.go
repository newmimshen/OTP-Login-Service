package handlers

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "otp-login-service/internal/storage"
)

// ✅ تعریف متغیرها با حرف بزرگ
var UserStore *storage.UserStorage
var JwtSecret []byte

type OTPVerifyRequest struct {
    Phone string `json:"phone" binding:"required"`
    Code  string `json:"code" binding:"required"`
}

// @Summary Verify OTP and login/register user
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body OTPVerifyRequest true "OTP verification payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /auth/verify-otp [post]
func VerifyOTPHandler(c *gin.Context) {
    var req OTPVerifyRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "شماره تلفن و کد OTP الزامی هستند"})
        return
    }

    entry, exists := storage.GetOTP(req.Phone)
    if !exists || time.Now().After(entry.ExpiresAt) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "OTP منقضی شده یا وجود ندارد"})
        return
    }

    if entry.Code != req.Code {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "OTP اشتباه است"})
        return
    }

    storage.DeleteOTP(req.Phone)

    if !UserStore.Exists(req.Phone) {
        UserStore.Add(req.Phone)
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "phone": req.Phone,
        "exp":   time.Now().Add(24 * time.Hour).Unix(),
    })

    tokenString, err := token.SignedString(JwtSecret)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در تولید توکن"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "ورود موفق",
        "token":   tokenString,
    })
}
