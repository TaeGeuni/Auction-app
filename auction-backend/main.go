package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:1504@tcp(127.0.0.1:3306)/auction_app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB 연결 실패: " + err.Error())
	}

	if err := RunMigrations(db); err != nil {
		panic("마이그레이션 실패: " + err.Error())
	}

	if err := Seed(db); err != nil {
		panic("시드 데이터 삽입 실패: " + err.Error())
	}

	r := gin.Default()
	r.POST("/api/auth/login", HandleLogin(db))

	fmt.Println("서버 시작: http://localhost:8080")
	r.Run(":8080")
}
