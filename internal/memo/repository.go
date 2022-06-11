package memo

import (
	"memo/internal/user"
)

type Repository interface {
	Create(memo *Memo) error
	Rand(qty uint, user *user.User) ([]Memo, error)
	Search(text string, user *user.User, limit uint) ([]Memo, error)
}
