package registry

import (
	// "github.com/manakuro/golang-clean-architecture/usecase/interactor"

	"github.com/nuoinguyen/gin-rubenv/interface/controller"
	"github.com/nuoinguyen/gin-rubenv/usecase"

	ip "github.com/nuoinguyen/gin-rubenv/interface/presenter"
	ir "github.com/nuoinguyen/gin-rubenv/interface/repository"
	// ur "github.com/nuoinguyen/gin-rubenv/usecase/repository"
)

func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() usecase.UserInteractor {
	return usecase.NewUserInteractor(r.NewUserRepository())
}

func (r *registry) NewUserRepository() ir.UserRepository {
	return ir.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() ip.UserPresenter {
	return ip.NewUserPresenter()
}
