package domain

//go:generate mockery --case=underscore --inpackage --name=UserRepository --filename=user_repository_mock.go --structname=UserRepositoryMock

type UserRepository interface {
	Create(user *User) error
	Save(user *User) error
	FindAll() ([]User, error)
	FindByTgAccount(tgAccount string) (*User, error)
}
