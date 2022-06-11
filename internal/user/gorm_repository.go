package user

import (
	"gorm.io/gorm"
)

type GORMRepository struct {
	db *gorm.DB
}

func (r *GORMRepository) Create(user *User) error {
	tx := r.db.Create(user)

	return tx.Error
}

func (r *GORMRepository) Save(user *User) error {
	tx := r.db.Save(user)

	return tx.Error
}

func (r *GORMRepository) FindAll() ([]User, error) {
	var users []User

	tx := r.db.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return users, nil
}

func (r *GORMRepository) FindByTgAccount(tgAccount string) (*User, error) {
	user := &User{}
	tx := r.db.Find(&user, "tg_account = ?", tgAccount)

	if tx.Error != nil {
		return nil, tx.Error
	}

	if user.ID == 0 {
		return nil, nil //nolint: nilnil
	}

	return user, nil
}

func NewGORMRepository(db *gorm.DB) *GORMRepository {
	if db == nil {
		panic("need db")
	}

	return &GORMRepository{
		db: db,
	}
}
