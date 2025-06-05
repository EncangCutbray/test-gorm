package models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Username string      `gorm:"size:100"`
	Email    string      `gorm:"size:255;not null;unique"`
	Password string      `gorm:"size:255"`
	Age      int64       `gorm:"default:0"`
	Note     []NoteModel `gorm:"foreignKey:UserID;references:ID"`
	Credit   CreditModel `gorm:"foreignKey:UserID;references:ID"`
}

func (UserModel) TableName() string {
	return "users"
}
