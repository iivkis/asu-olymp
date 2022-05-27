package repository

import (
	"gorm.io/gorm"
)

type AnswerModel struct {
	ID uint `gorm:"index:,unique" json:"id"`

	Value string `gorm:"size:1000" json:"value"`

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

func (r *AnswersRepository) Update(where *AnswerModel, fields map[string]interface{}) error {
	return r.db.Model(&AnswerModel{}).Where(where).Updates(fields).Error
}

func (r *AnswersRepository) Exists(where *AnswerModel) bool {
	return r.db.Select("id").First(&AnswerModel{}, where).Error == nil
}

func (r *AnswersRepository) FindAndTransformToMapByQuestionID(QuestionsID []uint) (m map[uint]*AnswerModel, err error) {
	m = make(map[uint]*AnswerModel)

	var answers []*AnswerModel
	if err = r.db.Where("question_id IN ?", QuestionsID).Find(&answers).Error; err != nil {
		return
	}

	for i := range answers {
		m[answers[i].QuestionID] = answers[i]
	}

	return
}
