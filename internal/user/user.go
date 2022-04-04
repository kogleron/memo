package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	TgAccount string
	TgChatID  int64
}
