package memo

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"memo/internal/user"
)

type Repository struct {
	db *gorm.DB
}

func (r *Repository) Create(memo *Memo) {
	r.db.Create(memo)
}

func (r *Repository) Rand(qty uint, user *user.User) []Memo {
	var memos []Memo

	if user == nil {
		return memos
	}

	r.db.
		Where("user_id = ?", user.ID).
		Clauses(
			clause.OrderBy{
				Expression: clause.Expr{SQL: "RANDOM()", WithoutParentheses: true},
			},
		).
		Limit(int(qty)).
		Find(&memos)

	return memos
}

func NewRepository(db *gorm.DB) *Repository {
	if db == nil {
		panic("need db")
	}

	return &Repository{
		db: db,
	}
}
