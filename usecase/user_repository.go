package usecase

import "github.com/nuoinguyen/gin-rubenv/domain/model"

// UserRepository ...
type UserRepository interface {
	// FindAll(u []*model.User) ([]*model.User, error)
	Store(user model.User) (us model.User, err error)
	FindAll() (event []*model.User, err error)
	FindById(id int) (event model.User, err error)
}
