package memo

import (
	"memo/internal/user"
)

type Repository interface {
	Create(memo *Memo) error
	Rand(qty uint, user *user.User) ([]Memo, error)
	Search(text string, user *user.User, limit uint) ([]Memo, error)
	FindByID(user *user.User, id uint) (*Memo, error)
	Delete(memo *Memo) error
}
