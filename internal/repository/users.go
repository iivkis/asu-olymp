package repository

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserModel struct {
	ID uint `gorm:"index:,unique"`

	Email    string
	FullName string
	Password string

	IsStudent bool
}

func (m *UserModel) BeforeCreate(tx *gorm.DB) error {
	b, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	m.Password = string(b)
	return nil
}

func (m *UserModel) BeforeUpdate(tx *gorm.DB) error {
	if m.Password != "" {
		b, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.MinCost)
		if err != nil {
			return err
		}
		m.Password = string(b)
	}

	return nil
}

type users struct {
	db *gorm.DB
}

func (r *users) Create(model *UserModel) error {
	return r.db.Create(model).Error
}

func (r *users) First(where *UserModel) (model *UserModel, err error) {
	err = r.db.Where(where).First(model).Error
	return
}

func (r *users) Exists(where *UserModel) bool {
	return r.db.Select("id").Where(where).First(&UserModel{}).Error == nil
}
