package domain

type MemoRepository interface {
	Create(memo *Memo) error
	Rand(qty uint, user *User) ([]Memo, error)
	Search(text string, user *User, limit uint) ([]Memo, error)
	FindByID(user *User, id uint) (*Memo, error)
	Delete(memo *Memo) error
}
