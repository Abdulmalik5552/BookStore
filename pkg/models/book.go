package models

import (
	"github.com/TheBusBoys/Tracker/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct{
	gorm.Model
	Name string
	Author string
	Publication string
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book{
	// db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) *Book{
	var getBook Book
	r := db.Where("ID = ?", Id).First(&getBook)
	if r.RowsAffected != 0 {
		return &getBook
	} else {
		return nil
	}
}

func DeleteBook(ID int64) *Book{
	var book Book
	r := db.Where("ID=?", ID).Delete(&book)
	if r.RowsAffected != 0 {
		return &book
	} else {
		return nil
	}
}