package database

import "github.com/devjoemedia/chitodopostgress/models"

func Migrate() {
	DB.AutoMigrate(
		&models.User{},
		&models.Todo{},
		&models.Ticket{},
		&models.RefreshToken{},
	)
}
