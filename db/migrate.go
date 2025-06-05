package db

import (
	"cutbray/test-gorm/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.UserModel{})
	db.AutoMigrate(&models.NoteModel{})
	db.AutoMigrate(&models.CreditModel{})
}
