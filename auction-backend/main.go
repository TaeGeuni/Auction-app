package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:Btg15041013?@tcp(127.0.0.1:3306)/auction_app?charset=utf8mb4&parseTime=True&loc=Local"
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

	fmt.Println("시드 데이터 삽입 성공!")
}
