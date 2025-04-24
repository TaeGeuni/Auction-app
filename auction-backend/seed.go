package main

import (
	"fmt"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	users := []User{
		{Email: "admin@example.com", PasswordHash: "hashed-password", Nickname: "관리자", Balance: 1000000, IsVerified: true},
		{Email: "user1@example.com", PasswordHash: "hashed-password", Nickname: "사용자1", Balance: 500000, IsVerified: true},
		{Email: "user2@example.com", PasswordHash: "hashed-password", Nickname: "사용자2", Balance: 300000, IsVerified: false},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			return fmt.Errorf("샘플 유저 삽입 실패: %w", err)
		}
	}

	return nil
}
