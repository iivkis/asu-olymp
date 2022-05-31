package repository

import (
	"time"

	"gorm.io/gorm"
)

type CheckingAnswerModel struct {
	ID uint `gorm:"index:,unique" json:"id"`

	PercentOfCorret uint  `json:"percent_of_correct"`
	CreatedAt       int64 `json:"created_at"`

	TaskID    uint      `json:"task_id"`
	TaskModel TaskModel `gorm:"foreignKey:TaskID" json:"-"`
}

func (m *CheckingAnswerModel) BeforeCreate(tx *gorm.DB) error {
	m.CreatedAt = time.Now().UTC().UnixMilli()
	return nil
}

type CheckingAnswersRepository struct {
	db *gorm.DB
}

func NewCheckingAnswersRepository(db *gorm.DB) *CheckingAnswersRepository {
	return &CheckingAnswersRepository{
		db: db,
	}
}

func (r *CheckingAnswersRepository) Cursor() *gorm.DB {
	return r.db.Model(&CheckingAnswersRepository{})
}
