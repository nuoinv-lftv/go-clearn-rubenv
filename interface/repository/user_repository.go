package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/nuoinguyen/gin-rubenv/domain/model"
)

type userRepository struct {
	db *gorm.DB
}

// UserRepository ..
type UserRepository interface {
	FindAll() (users []*model.User, err error)
	Store(data model.User) (user model.User, err error)
	FindById(id int) (user model.User, err error)
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

func (repo *userRepository) Store(data model.User) (user model.User, err error) {
	// fmt.Println(data)
	// user = model.User{}
	repo.db.NewRecord(data)
	repo.db.Create(&data)

	return data, err
}

func (repo *userRepository) FindById(id int) (user model.User, err error) {
	user = model.User{}
	repo.db.Where("ID = ?", id).First(&user)
	if user.ID <= 0 {
		return model.User{}, errors.New("User, not found")
	}
	return user, nil
}
