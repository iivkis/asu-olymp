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

func (r *UsersRepository) Create(model *UserModel) error {
	return r.db.Create(model).Error
}

func (r *UsersRepository) First(where *UserModel) (model *UserModel, err error) {
	err = r.db.Where(where).First(&model).Error
	return
}

func (r *UsersRepository) Exists(where *UserModel) bool {
	return r.db.Select("id").Where(where).First(&UserModel{}).Error == nil
}

func (r *UsersRepository) SignUpByEmail(email string, password string) (model *UserModel, err error) {
	user, err := r.First(&UserModel{Email: email})
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}
