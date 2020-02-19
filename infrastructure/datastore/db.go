package datastore

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/nuoinguyen/gin-rubenv/config"
)

// NewDB ..
func NewDB() *gorm.DB {
	DBMS := "mysql"
	mySqlConfig := &mysql.Config{
		User:                 config.C.Database.User,
		Passwd:               config.C.Database.Password,
		Net:                  config.C.Database.Net,
		Addr:                 config.C.Database.Addr,
		DBName:               config.C.Database.DBName,
		AllowNativePasswords: config.C.Database.AllowNativePasswords,
		Params: map[string]string{
			"parseTime": config.C.Database.Params.ParseTime,
		},
	}

	db, err := gorm.Open(DBMS, mySqlConfig.FormatDSN())

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

type DB interface {
	Create(value interface{}) *gorm.DB
	NewRecord(value interface{}) bool
	Find(out interface{}, where ...interface{}) *gorm.DB
	First(out interface{}, where ...interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
}
