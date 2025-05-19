package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var jwtKey = []byte("secret-key") // 진짜 배포에서는 환경변수로 관리해!

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func HandleLogin(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 요청"})
			return
		}

		var user User
		if err := db.First(&user, "email = ?", req.Email).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "이메일 또는 비밀번호가 틀렸습니다"})
			return
		}

		// 임시: 비밀번호 비교는 bcrypt가 아니라 간단히 문자열 비교
		if req.Password != user.PasswordHash {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "이메일 또는 비밀번호가 틀렸습니다"})
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": user.ID,
			"exp":     time.Now().Add(24 * time.Hour).Unix(),
		})

		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "토큰 생성 실패"})
			return
		}

		c.JSON(http.StatusOK, LoginResponse{Token: tokenString})
	}
}
