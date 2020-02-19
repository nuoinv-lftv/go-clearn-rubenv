package usecase

import "github.com/nuoinguyen/gin-rubenv/domain/model"

// UserPresenter ..
type UserPresenter interface {
	ResponseUser(u []*model.User) []*model.User
}
