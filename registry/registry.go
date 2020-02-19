package registry

import (
	"github.com/jinzhu/gorm"
	"github.com/nuoinguyen/gin-rubenv/interface/controller"
)

type registry struct {
	db *gorm.DB
}

// Registry ...
type Registry interface {
	NewAppController() controller.AppController
}

// NewRegistry ...
func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

// NewAppController ..
func (r *registry) NewAppController() controller.AppController {
	return r.NewUserController()
}

// // NewUserController ..
// func (i *registry) NewUserController() controller.AppController {
// 	return controller.NewUserController(i.NewUserInteractor())
// }

// // NewUserInteractor ..
// func (i *registry) NewUserInteractor() usecase.UserInteractor {
// 	return usecase.NewUserInteractor(i.NewUsersRepository())
// }

// // NewUsersRepository ..
// func (i *registry) NewUsersRepository() usecase.UserRepository {
// 	return usecase.NewUserRepository(i.db)
// }
