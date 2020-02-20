package usecase

import (
	"github.com/nuoinguyen/gin-rubenv/domain/model"
)

type userInteractor struct {
	User UserRepository
	// pres UserPresenter // c1
}

// UserInteractor ...
type UserInteractor interface {
	GetAll() ([]*model.User, error)
	AddUser(data model.User) (user model.User, err error)
	UserById(id int) (user model.User, err error)
}

// NewUserInteractor ...
// func NewUserInteractor(u UserRepository, p UserPresenter) UserInteractor {
func NewUserInteractor(u UserRepository) UserInteractor { // C1
	// return &userInteractor{u, p} // c1
	return &userInteractor{User: u}
}

// GetAll ...
func (us *userInteractor) GetAll() (u []*model.User, err error) {
	u, err = us.User.FindAll()

	if err != nil {
		return []*model.User{}, err
	}

	// return us.pres.ResponseUsers(u), nil // c1
	return u, nil
}

// AddUser ...
func (interactor *userInteractor) AddUser(data model.User) (user model.User, err error) {
	user, err = interactor.User.Store(data)
	if err != nil {
		return model.User{}, err
	}
	return user, err
}

// UserById ...
func (interactor *userInteractor) UserById(id int) (user model.User, err error) {
	user, err = interactor.User.FindById(id)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
