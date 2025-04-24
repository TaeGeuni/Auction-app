package main

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Email        string `gorm:"unique"`
	PasswordHash string
	Nickname     string
	Balance      int64
	IsVerified   bool
}

func RunMigrations(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "20250424_create_users",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&User{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("users")
			},
		},
	})

	return m.Migrate()
}
