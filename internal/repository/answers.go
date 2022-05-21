package repository

import "gorm.io/gorm"

type AnswerModel struct {
	ID uint `gorm:"index:,unique" json:"id"`

	Value string `gorm:"size:1000" json:"value,omitempty"`

	QuestionID uint `json:"question_id"`
	AuthorID   uint `json:"-"`

	QuestionModel QuestionModel `gorm:"foreignKey:QuestionID" json:"-"`
	UserModel     UserModel     `gorm:"foreignKey:AuthorID" json:"-"`
}

type AnswersRepository struct {
	db *gorm.DB
}

func NewAnswersRepository(db *gorm.DB) *AnswersRepository {
	return &AnswersRepository{db: db}
}

func (r *AnswersRepository) Cursor() *gorm.DB {
	return r.db.Model(&AnswerModel{})
}

func (r *AnswersRepository) Find(where *AnswerModel, payload *Payload) (models []*AnswerModel, err error) {
	err = r.db.Where("id > ?", payload.OffsetID).Limit(payload.Limit).Find(&models, where).Error
	return
}
