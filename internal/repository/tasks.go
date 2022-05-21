package repository

import "gorm.io/gorm"

type TaskModel struct {
	ID       uint   `gorm:"index:,unique"`
	Title    string `gorm:"size:200"`
	Content  string `gorm:"size:2000"`
	AuthorID uint

	UserModel UserModel `gorm:"foreignKey:AuthorID"`
}

type TasksRepository struct {
	db *gorm.DB
}

func NewTasksRepository(db *gorm.DB) *TasksRepository {
	return &TasksRepository{db: db}
}

func (r *TasksRepository) CreateTask(model *TaskModel) error {
	return r.db.Create(model).Error
}

func (r *TasksRepository) Update(where *TaskModel, fields map[string]interface{}) error {
	return r.db.Model(&TaskModel{}).Where(where).Updates(fields).Error
}
