package repository

import "gorm.io/gorm"

type TaskModel struct {
	ID      uint   `gorm:"index:,unique"`
	Title   string `gorm:"varchar(200)"`
	Content string `gorm:"varchar(2000)"`
}

type TasksRepository struct {
	db *gorm.DB
}

func NewTasksRepository(db *gorm.DB) *TasksRepository {
	return &TasksRepository{db: db}
}

func (r *TasksRepository) CreateTask(model *TasksRepository) error {
	return r.db.Create(model).Error
}

func (r *TasksRepository) Update(where *TaskModel, model interface{}) error {
	return r.db.Model(&TaskModel{}).Where(where).Updates(model).Error
}
