package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/nuoinguyen/gin-rubenv/domain/model"
)

type userRepository struct {
	db *gorm.DB
}

// UserRepository ..
type UserRepository interface {
	// FindAll() (u []*model.User, error)
	FindAll() (users []*model.User, err error)
}

// NewUserRepository ..
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (repo *userRepository) FindAll() (users []*model.User, err error) {
	// err := repo.db.Find(&u).Error

	// if err != nil {
	// 	return nil, err
	// }

	// return u, nil

	users = []*model.User{}
	repo.db.Find(&users)
	return users, err
}
