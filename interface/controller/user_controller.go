package controller

import (

	// "github.com/gin-gonic/gin"

	"net/http"

	"github.com/nuoinguyen/gin-rubenv/usecase"
)

type userController struct {
	usecase usecase.UserInteractor
}

// UserController ..
type UserController interface {
	GetUsers(c Context) error
}

// NewUserController ..
func NewUserController(us usecase.UserInteractor) UserController {
	return &userController{us}
}

func (uc *userController) GetUsers(c Context) error {
	// var u []*model.User

	u, err := uc.usecase.GetAll()
	// u, err := uc.userInteractor.Get(u)
	if err != nil {
		return err
	}

	// res := Response{Message: u}

	return c.JSON(http.StatusOK, u)
}
