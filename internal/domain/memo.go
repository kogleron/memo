package domain

import (
	"gorm.io/gorm"
)

type Memo struct {
	gorm.Model
	Text   string
	UserID uint
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
