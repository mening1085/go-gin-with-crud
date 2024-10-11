package middleware

import (
	"go-crud/utils" // นำเข้าฟังก์ชัน ValidateToken
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware ตรวจสอบ Bearer token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ดึง Authorization header
		token := c.Request.Header.Get("Authorization")

		// ตรวจสอบว่า token เป็น Bearer token
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No or invalid authorization header"})
			c.Abort()
			return
		}

		// เอา token ออกจาก "Bearer " prefix
		tokenString := strings.TrimPrefix(token, "Bearer ")

		// ตรวจสอบความถูกต้องของ token
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// ตั้งค่า username ใน context สำหรับใช้งานใน handler ถัดไป
		c.Set("username", claims.Username)
		c.Next()
	}
}
