package database

import "github.com/devjoemedia/chitodopostgress/models"

func Migrate() {
	DB.AutoMigrate(
		&models.Todo{},
		&models.User{},
		&models.RefreshToken{},
	)
}
