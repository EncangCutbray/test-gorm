package models

import "gorm.io/gorm"

type NoteModel struct {
	gorm.Model
	Content string     `gorm:"size:255"`
	UserID  uint       `gorm:"index"`
	User    *UserModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Belongs To
}

func (NoteModel) TableName() string {
	return "notes"
}
