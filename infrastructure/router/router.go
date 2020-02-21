package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/nuoinguyen/gin-rubenv/interface/controller"
)

type Routing interface {
	NewRouter(ctrl controller.UserController)
	Run(port string)
}

type routing struct {
	Gin    *gin.Engine
	Router *mux.Router
}

func NewRouting() Routing {
	return &routing{
		Gin: gin.Default(),
	}
}

// func (r *routing) SetRouting(ctrl controller.UserController) {
// 	u := r.Gin.Group("/api/v1/user")
// 	{
// 		// Create
// 		u.POST("/register", func(c *gin.Context) { ctrl.AddUser(c) })
// 		// Read
// 		u.GET("/get", func(c *gin.Context) { ctrl.GetUsers(c) })
// 		// u.GET("/get/:id", func(c *gin.Context) { ctrl.UserById(c) })
// 		// Update (skip)
// 		// Delete (skip)
// 	}
// }

// NewRouter ..
func (r *routing) NewRouter(ctrl controller.UserController) {
	r.Gin.Static("/static", "./static")

	u := r.Gin.Group("/api/v1/user")
	{
		u.POST("/register", func(c *gin.Context) { ctrl.AddUser(c) })
		u.GET("/get", func(c *gin.Context) { ctrl.GetUsers(c) })
		u.GET("/get/:id", func(c *gin.Context) { ctrl.UserById(c) })
	}
}

func (r *routing) Run(port string) {
	r.Gin.Run(port)

	fmt.Printf("%v", r)
}

// // NewRouter ..
// func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
// 	e.Use(middleware.Logger())
// 	e.Use(middleware.Recover())

// 	e.GET("/users", func(context echo.Context) error { return c.GetUsers(context) })
// 	// Create
// 	// e.POST("/register", func(c echo.Context) error { c.AddUser(context) })

// 	// e.GET("/get/:id", func(c echo.Context) error { close.UserById(context) })

// 	return e
// }
