package repository

import (
	"gorm.io/gorm"
)

type QuestionModel struct {
	ID   uint   `gorm:"index:,unique" json:"id"`
	Text string `json:"text"`

	TaskID   uint `json:"task_id"`
	AuthorID uint `json:"-"`

	TaskModel TaskModel `gorm:"foreignKey:TaskID" json:"-"`
	UserModel UserModel `gorm:"foreignKey:AuthorID" json:"-"`
}

type QuestionsRepository struct {
	db *gorm.DB
}

func NewQuestionsRepository(db *gorm.DB) *QuestionsRepository {
	return &QuestionsRepository{
		db: db,
	}
}

func (r *QuestionsRepository) Cursor() *gorm.DB {
	return r.db.Model(&QuestionModel{})
}

func (r *QuestionsRepository) Find(where *QuestionModel, payload *Payload) (models []*QuestionModel, err error) {
	err = r.db.Where("id > ?", payload.OffsetID).Limit(payload.Limit).Find(&models, where).Error
	return
}

func (r *QuestionsRepository) Update(where *QuestionModel, fields map[string]interface{}) error {
	return r.db.Model(&QuestionModel{}).Where(where).Updates(fields).Error
}

func (r *QuestionsRepository) Exists(where *QuestionModel) bool {
	return r.db.Select("id").First(&QuestionModel{}, where).Error == nil
}

func (r *QuestionsRepository) HashmapOfExistence(where *QuestionModel) (m map[uint]bool, err error) {
	m = make(map[uint]bool)

	var q []QuestionModel
	if err = r.db.Select("id").Find(&q, where).Error; err != nil {
		return
	}

	for i := range q {
		m[q[i].ID] = true
	}
	return
}
