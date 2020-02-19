package usecase

import (
	"github.com/nuoinguyen/gin-rubenv/domain/model"
)

type userInteractor struct {
	User UserRepository
	// pres UserPresenter
}

// UserInteractor ...
type UserInteractor interface {
	Get() ([]*model.User, error)
	// Add(data model.User) (err error)
	// Users() (users []model.User, err error)
	// UserById(id int) (user model.User, err error)
}

// NewUserInteractor ...
func NewUserInteractor(repo UserRepository) UserInteractor {
	return &userInteractor{User: repo}
}

// func NewUserInteractor(repo UserRepository) UserInteractor {
// 	return &userInteractor{User: repo}
// }

// Get ...
func (us *userInteractor) Get() (u []*model.User, err error) {
	u, err = us.User.FindAll()

	if err != nil {
		return []*model.User{}, err
	}

	// return us.UserPresenter.ResponseUser(u), nil
	return u, nil
}

// func (interactor *userInteractor) Users() (users []domain.Users, err error) {
// 	users, err = interactor.User.FindAll()
// 	if err != nil {
// 		return []domain.Users{}, err
// 	}
// 	return users, nil
// }
