package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Secret key สำหรับเข้ารหัส JWT
var jwtKey = []byte("your_secret_key")

// ข้อมูลใน Token
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// สร้าง Token
func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // ตั้งเวลาในการหมดอายุเป็น 24 ชั่วโมง
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ตรวจสอบ Token
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}
	return claims, nil
}
