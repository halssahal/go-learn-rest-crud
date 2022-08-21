package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func SetupDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/learn_go_crud?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}

	return db
}
