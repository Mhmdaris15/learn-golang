package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "root:@tcp(localhost:3306)/librarydb?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db = d
	fmt.Println("Connected to the database!", dsn)
}

func GetDB() *gorm.DB {
	return db
}
