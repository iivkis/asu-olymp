package repository

import "gorm.io/gorm"

type Repository struct {
	Users *UsersRepository
}

func NewRepository() (r *Repository) {
	//new connection
	db := newConnection()

	//migrations
	if err := migration(db); err != nil {
		panic(err)
	}

	//combine
	return &Repository{
		Users: NewUsersRepository(db),
	}
}

func migration(db *gorm.DB) error {
	return db.AutoMigrate(
		&UserModel{},
	)
}
