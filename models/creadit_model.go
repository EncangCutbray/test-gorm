package models

import "gorm.io/gorm"

type CreditModel struct {
	gorm.Model
	Number string     `gorm:"size:255"`
	UserID uint       `gorm:"index"`
	User   *UserModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Belongs To
}

func (CreditModel) TableName() string {
	return "credits"
}
