package handlers

import (
    "net/http"
    "strconv"
    "strings"

    "github.com/gin-gonic/gin"
    "otp-login-service/internal/storage"
)


// @Summary لیست کاربران با جستجو و صفحه‌بندی
// @Tags Users
// @Produce json
// @Param page query int false "شماره صفحه"
// @Param limit query int false "تعداد در هر صفحه"
// @Param search query string false "جستجو بر اساس شماره"
// @Success 200 {object} map[string]interface{}
// @Router /users [get]
func ListUsersHandler(c *gin.Context) {
    pageStr := c.DefaultQuery("page", "1")
    limitStr := c.DefaultQuery("limit", "10")
    search := c.Query("search")

    page, _ := strconv.Atoi(pageStr)
    limit, _ := strconv.Atoi(limitStr)
    if page < 1 {
        page = 1
    }

    allUsers := UserStore.List()

    filtered := []storage.User{}
    for _, u := range allUsers {
        if search == "" || strings.Contains(u.Phone, search) {
            filtered = append(filtered, u)
        }
    }

    start := (page - 1) * limit
    end := start + limit
    if start > len(filtered) {
        start = len(filtered)
    }
    if end > len(filtered) {
        end = len(filtered)
    }

    c.JSON(http.StatusOK, gin.H{
        "users": filtered[start:end],
        "total": len(filtered),
        "page":  page,
        "limit": limit,
    })
}

// @Summary دریافت اطلاعات یک کاربر خاص
// @Tags Users
// @Produce json
// @Param phone path string true "شماره تلفن"
// @Success 200 {object} storage.User
// @Failure 404 {object} map[string]string
// @Router /users/{phone} [get]
func GetUserByPhoneHandler(c *gin.Context) {
    phone := c.Param("phone")
    user, exists := UserStore.Get(phone)
    if !exists {
        c.JSON(http.StatusNotFound, gin.H{"error": "کاربر پیدا نشد"})
        return
    }

    c.JSON(http.StatusOK, user)
}
