package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func init() {

	dbTmp, err := gorm.Open(
		"postgres",
		"host=D8C74545-postgres port=5432 user=psystem dbname=psystem password=psystem sslmode=disable")
	if err != nil {
		panic(err)
	}
	db = dbTmp

}
