package infra

import (
	"gorm.io/gorm"

	"memo/internal/domain"
)

func NewUserGORMRepository(db *gorm.DB) *UserGORMRepository {
	if db == nil {
		panic("need db")
	}

	return &UserGORMRepository{
		db: db,
	}
}

type UserGORMRepository struct {
	db *gorm.DB
}

func (r *UserGORMRepository) Create(user *domain.User) error {
	tx := r.db.Create(user)

	return tx.Error
}

func (r *UserGORMRepository) Save(user *domain.User) error {
	tx := r.db.Save(user)

	return tx.Error
}

func (r *UserGORMRepository) FindAll() ([]domain.User, error) {
	var users []domain.User

	tx := r.db.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return users, nil
}

func (r *UserGORMRepository) FindByTgAccount(tgAccount string) (*domain.User, error) {
	user := &domain.User{}
	tx := r.db.Find(&user, "tg_account = ?", tgAccount)

	if tx.Error != nil {
		return nil, tx.Error
	}

	if user.ID == 0 {
		return nil, nil //nolint: nilnil
	}

	return user, nil
}
