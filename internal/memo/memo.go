package memo

import (
	"gorm.io/gorm"

	"memo/internal/user"
)

type Memo struct {
	gorm.Model
	Text   string
	UserID uint
	User   user.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
