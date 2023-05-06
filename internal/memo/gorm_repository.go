package memo

import (
	"errors"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"memo/internal/user"
)

var (
	errEmptyUser   = errors.New("empty user")
	errEmptyGORMDB = errors.New("empty gorm db")
)

type GORMRepository struct {
	db *gorm.DB
}

func (r *GORMRepository) Create(memo *Memo) error {
	tx := r.db.Create(memo)

	return tx.Error
}

func (r *GORMRepository) Rand(qty uint, user *user.User) ([]Memo, error) {
	var memos []Memo

	if user == nil {
		return nil, errEmptyUser
	}

	tx := r.db.
		Where("user_id = ?", user.ID).
		Clauses(
			clause.OrderBy{
				Expression: clause.Expr{SQL: "RANDOM()", WithoutParentheses: true},
			},
		).
		Limit(int(qty)).
		Find(&memos)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return memos, nil
}

func (r *GORMRepository) Search(text string, user *user.User, limit uint) ([]Memo, error) {
	var memos []Memo

	if user == nil {
		return nil, errEmptyUser
	}

	prepText := strings.Trim(text, " ")
	if len(prepText) == 0 {
		return memos, nil
	}

	tx := r.db.
		Where("user_id = ? AND text LIKE ?", user.ID, "%"+prepText+"%").
		Limit(int(limit)).
		Find(&memos)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return memos, nil
}

func (r *GORMRepository) FindByID(user *user.User, id uint) (*Memo, error) {
	var memos []Memo

	if user == nil {
		return nil, errEmptyUser
	}

	tx := r.db.
		Where("user_id = ? AND id = ?", user.ID, id).
		Limit(1).
		Find(&memos)

	if tx.Error != nil {
		return nil, tx.Error
	}

	if len(memos) == 0 {
		return nil, nil //nolint
	}

	return &memos[0], nil
}

func (r *GORMRepository) Delete(memo *Memo) error {
	return r.db.Delete(memo).Error
}

func NewGORMRepository(db *gorm.DB) (*GORMRepository, error) {
	if db == nil {
		return nil, errEmptyGORMDB
	}

	return &GORMRepository{
		db: db,
	}, nil
}
