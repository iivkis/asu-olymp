package repository

import "gorm.io/gorm"

type TaskModel struct {
	ID uint `gorm:"index:,unique" json:"id"`

	Title   string `gorm:"size:200" json:"title"`
	Content string `gorm:"size:2000" json:"content"`

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
	ID uint `gorm:"index:,unique" json:"id"`

	Title   string `gorm:"size:200" json:"title"`
	Content string `gorm:"size:2000" json:"content"`

	SolutionsCount uint `json:"solutions_count"`

	AuthorID uint   `json:"author_id"`
	FullName string `json:"author_name"`
}

func (r *TasksRepository) Find(where *TaskModel, payload *Payload) (models []*TasksFindResult, err error) {
	err = r.Cursor().
		Select("task_models.id, task_models.title, task_models.content, task_models.solutions_count, task_models.author_id, user_models.full_name").
		Joins("JOIN user_models ON user_models.id = task_models.author_id").
		Where("task_models.id > ?", payload.OffsetID).
		Limit(payload.Limit).
		Find(&models, where).Error
	return
}

func (r *TasksRepository) FindByID(id uint) (models *TasksFindResult, err error) {
	err = r.Cursor().
		Select("task_models.id, task_models.title, task_models.content, task_models.solutions_count, task_models.author_id, user_models.full_name").
		Joins("JOIN user_models ON user_models.id = task_models.author_id").
		Where("task_models.id = ?", id).
		Find(&models).Error
	return
}

func (r *TasksRepository) Update(where *TaskModel, fields map[string]interface{}) error {
	return r.db.Model(&TaskModel{}).Where(where).Updates(fields).Error
}

func (r *TasksRepository) Exists(where *TaskModel) bool {
	return r.db.Select("id").First(&TaskModel{}, where).Error == nil
}
