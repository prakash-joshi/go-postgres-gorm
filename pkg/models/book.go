package models

import (
	"github.com/jinzhu/gorm"
	"github.com/prakash-joshi/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name         string `json:"name"`
	Author       string `json:"author"`
	Publications string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("Id=?", Id).Find(&getBook)
	return &getBook, db
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("Id=?", Id).Delete(book)
	return book
}
