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

type UsersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

func (c *UsersRepository) Cursor() *gorm.DB {
	return c.db.Model(&UserModel{})
}

func (r *UsersRepository) Exists(where *UserModel) bool {
	return r.db.Select("id").Where(where).First(&UserModel{}).Error == nil
}

func (r *UsersRepository) SignInByEmail(email string, password string) (model *UserModel, err error) {
	if err = r.db.First(&model, &UserModel{Email: email}).Error; err != nil {
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(model.Password), []byte(password))
	return
}
