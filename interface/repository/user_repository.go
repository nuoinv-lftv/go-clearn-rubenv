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
	FindAll() (users []*model.User, err error)
}

// NewUserRepository ..
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (repo *userRepository) FindAll() (users []*model.User, err error) {
	users = []*model.User{}
	repo.db.Find(&users)

	return users, err
}
