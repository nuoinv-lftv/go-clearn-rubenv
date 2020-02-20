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
