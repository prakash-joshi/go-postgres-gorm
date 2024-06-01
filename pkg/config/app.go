package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "postgres://postgres:postgres@localhost:5432/bookstore?sslmode=disable"
	d, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDb() *gorm.DB {
	return db
}
