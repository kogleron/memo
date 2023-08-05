package infra

import (
	"errors"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"memo/internal/domain"
)

var (
	errEmptyUser   = errors.New("empty user")
	errEmptyGORMDB = errors.New("empty gorm db")
)

func NewMemoGORMRepository(db *gorm.DB) (*MemoGORMRepository, error) {
	if db == nil {
		return nil, errEmptyGORMDB
	}

	return &MemoGORMRepository{
		db: db,
	}, nil
}

type MemoGORMRepository struct {
	db *gorm.DB
}

func (r *MemoGORMRepository) Create(memo *domain.Memo) error {
	tx := r.db.Create(memo)

	return tx.Error
}

func (r *MemoGORMRepository) Rand(qty uint, user *domain.User) ([]domain.Memo, error) {
	var memos []domain.Memo

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

func (r *MemoGORMRepository) Search(text string, user *domain.User, limit uint) ([]domain.Memo, error) {
	var memos []domain.Memo

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

func (r *MemoGORMRepository) FindByID(user *domain.User, id uint) (*domain.Memo, error) {
	var memos []domain.Memo

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

func (r *MemoGORMRepository) Delete(memo *domain.Memo) error {
	return r.db.Delete(memo).Error
}
