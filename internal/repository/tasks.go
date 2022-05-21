package repository

import "gorm.io/gorm"

type TaskModel struct {
	ID uint `gorm:"index:,unique" json:"id"`

	Title    string `gorm:"size:200" json:"title"`
	Content  string `gorm:"size:2000" json:"content"`
	AuthorID uint   `json:"author_id"`

	UserModel UserModel `gorm:"foreignKey:AuthorID" json:"-"`
}

type TasksRepository struct {
	db *gorm.DB
}

func NewTasksRepository(db *gorm.DB) *TasksRepository {
	return &TasksRepository{db: db}
}

func (r *TasksRepository) Cursor() *gorm.DB {
	return r.db.Model(&TaskModel{})
}

func (r *TasksRepository) Find(where *TaskModel, payload *Payload) (models []*TaskModel, err error) {
	err = r.db.Where("id > ?", payload.OffsetID).Limit(payload.Limit).Find(&models, where).Error
	return
}

func (r *TasksRepository) Update(where *TaskModel, fields map[string]interface{}) error {
	return r.db.Model(&TaskModel{}).Where(where).Updates(fields).Error
}

func (r *TasksRepository) Exists(where *TaskModel) bool {
	return r.db.Select("id").First(&TaskModel{}, where).Error == nil
}
