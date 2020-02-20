package registry

import (
	"github.com/nuoinguyen/gin-rubenv/interface/controller"
	"github.com/nuoinguyen/gin-rubenv/usecase"

	ip "github.com/nuoinguyen/gin-rubenv/interface/presenter"
	ir "github.com/nuoinguyen/gin-rubenv/interface/repository"
	// ur "github.com/nuoinguyen/gin-rubenv/usecase/repository"
	// up "github.com/nuoinguyen/gin-rubenv/usecase/presenter"
)

func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() usecase.UserInteractor {
	// return usecase.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter()) // c1
	return usecase.NewUserInteractor(r.NewUserRepository())
}

func (r *registry) NewUserRepository() ir.UserRepository {
	return ir.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() usecase.UserPresenter {
	return ip.NewUserPresenter()
}
