package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/nuoinguyen/gin-rubenv/config"
	"github.com/nuoinguyen/gin-rubenv/infrastructure/datastore"
	"github.com/nuoinguyen/gin-rubenv/infrastructure/router"
	"github.com/nuoinguyen/gin-rubenv/registry"
)

func main() {
	config.ReadConfig()

	db := datastore.NewDB()
	db.LogMode(true)
	defer db.Close()

	r := registry.NewRegistry(db)

	ctrl := r.NewAppController()

	// e := echo.New()

	// debug mode
	// e.Debug = true

	e := router.NewRouting()
	// e = router.NewRouting()
	e.NewRouter(ctrl)

	// e.Debug = true

	e.Run(":8080")
	// fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	// if err := e.Start(":" + config.C.Server.Address); err != nil {
	// 	log.Fatalln(err)
	// }
}
