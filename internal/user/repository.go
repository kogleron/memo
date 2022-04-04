package user

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func (r *Repository) Create(user *User) {
	r.db.Create(user)
}

func (r *Repository) Save(user *User) {
	r.db.Save(user)
}

func (r *Repository) FindAll() []User {
	var users []User

	r.db.Find(&users)

	return users
}

func (r *Repository) FindByTgAccount(tgAccount string) *User {
	user := &User{}
	r.db.Find(&user, "tg_account = ?", tgAccount)

	if user.ID == 0 {
		return nil
	}

	return user
}

func NewRepository(db *gorm.DB) *Repository {
	if db == nil {
		panic("need db")
	}

	return &Repository{
		db: db,
	}
}
