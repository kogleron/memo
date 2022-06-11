package user

type Repository interface {
	Create(user *User) error
	Save(user *User) error
	FindAll() ([]User, error)
	FindByTgAccount(tgAccount string) (*User, error)
}
