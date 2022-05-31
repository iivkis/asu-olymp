package repository

import "gorm.io/gorm"

type TaskModel struct {
	ID uint `gorm:"index:,unique" json:"id"`

	Title   string `gorm:"size:200" json:"title"`
	Content string `gorm:"size:2000" json:"content"`

	ShowCorrect bool `json:"show_correct"`
	IsPublic    bool `json:"is_public"`

	SolutionsCount uint `json:"solutions_count"`

	AuthorID uint `json:"author_id"`

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

type TasksFindResult struct {
	ID uint `json:"id"`

	Title   string `json:"title"`
	Content string `json:"content"`

	ShowCorrect bool `json:"show_correct"`
	IsPublic    bool `json:"is_public"`

	SolutionsCount uint `json:"solutions_count"`

	AuthorID uint   `json:"author_id"`
	FullName string `json:"author_name"`
}

func (r *TasksRepository) Find(where *TaskModel, payload *Payload) (models []*TasksFindResult, err error) {
	err = r.db.Table("task_models AS tm").
		Select("tm.id, tm.title, tm.content, tm.show_correct, tm.is_public, tm.solutions_count, tm.author_id, um.full_name").
		Joins("JOIN user_models AS um ON um.id = tm.author_id").
		Where("tm.id > ?", payload.OffsetID).
		Order("tm.id DESC").
		Limit(payload.Limit).
		Find(&models, where).Error
	return
}

func (r *TasksRepository) FindByID(id uint) (models *TasksFindResult, err error) {
	err = r.Cursor().Table("task_models AS tm").
		Select("tm.id, tm.title, tm.content, tm.show_correct, tm.is_public, tm.solutions_count, tm.author_id, um.full_name").
		Joins("JOIN user_models AS um ON um.id = tm.author_id").
		Where("tm.id = ?", id).
		First(&models).Error
	return
}

func (r *TasksRepository) Update(where *TaskModel, fields map[string]interface{}) error {
	return r.db.Model(&TaskModel{}).Where(where).Updates(fields).Error
}

func (r *TasksRepository) Exists(where *TaskModel) bool {
	return r.db.Select("id").First(&TaskModel{}, where).Error == nil
}
