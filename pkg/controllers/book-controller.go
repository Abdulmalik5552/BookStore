package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/TheBusBoys/Tracker/pkg/models"
	"github.com/TheBusBoys/Tracker/pkg/utils"

	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	if len(books) != 0 {
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No books found"))
	}
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	if bookDetails != nil {
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Book with ID=" + bookId + " not found"))
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	if book != nil {
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Book with ID=" + bookId + " not found"))
	}
}

// This function should be in the model and not here
// func UpdateBook(w http.ResponseWriter, r *http.Request){
// 	var updateBook = &models.Book{}
// 	utils.ParseBody(r, updateBook)
// 	vars := mux.Vars(r)
// 	bookId := vars["bookId"]
// 	ID, err := strconv.ParseInt(bookId, 0,0)
// 	if err != nil {
// 		fmt.Println("error while parsing")
// 	}
// 	bookDetails, db:=models.GetBookById(ID)
// 	if updateBook.Name != ""{
// 		bookDetails.Name = updateBook.Name
// 	}
// 	if updateBook.Author != ""{
// 		bookDetails.Author = updateBook.Author
// 	}
// 	if updateBook.Publication != ""{
// 		bookDetails.Publication = updateBook.Publication
// 	}
// 	db.Save(&bookDetails)
// 	res, _ := json.Marshal(bookDetails)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }
