package memo

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func (r *Repository) Create(memo *Memo) {
	r.db.Create(memo)
}

func NewRepository(db *gorm.DB) *Repository {
	if db == nil {
		panic("need db")
	}

	return &Repository{
		db: db,
	}
}
