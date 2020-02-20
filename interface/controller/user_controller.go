package controller

import (

	// "github.com/gin-gonic/gin"

	"strconv"

	"github.com/nuoinguyen/gin-rubenv/domain/model"
	"github.com/nuoinguyen/gin-rubenv/usecase"
)

type userController struct {
	usecase usecase.UserInteractor
}

// UserController ..
type UserController interface {
	AddUser(c Context)
	GetUsers(c Context)
	UserById(c Context)
}

// NewUserController ..
func NewUserController(us usecase.UserInteractor) UserController {
	return &userController{us}
}

// AddUser ...
func (ctrl *userController) AddUser(c Context) {
	data := model.User{}

	// dump obj
	// spew.Dump(data)

	if err := c.BindJSON(&data); err != nil {
		c.String(500, "Bad request: "+err.Error())
	}

	user, err := ctrl.usecase.AddUser(data)

	// if err := ctrl.usecase.AddUser(data); err != nil {
	// 	c.JSON(409, err)
	// }
	if err != nil {
		c.JSON(400, err)
		return
	}

	c.JSON(201, user)
}

// GetUsers ...
func (uc *userController) GetUsers(c Context) {
	// var u []*model.User

	u, err := uc.usecase.GetAll()
	// u, err := uc.userInteractor.Get(u)
	if err != nil {
		c.JSON(404, err)
		return
	}

	c.JSON(200, u)
	// return c.JSON(200, u)
}

// UserById ...
func (uc *userController) UserById(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	u, err := uc.usecase.UserById(id)

	if err != nil {
		c.JSON(400, err)
		return
	}

	c.JSON(200, u)

}
