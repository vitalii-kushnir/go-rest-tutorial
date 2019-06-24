package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"go-rest-tutorial/models"
	"go-rest-tutorial/repository/book"
	"log"
	"net/http"
	"strconv"
)

type BookController struct{}

func (c BookController) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Get all book")
		bookRepo := bookRepository.BookRepository{}
		books := bookRepo.GetBooks(db)
		json.NewEncoder(w).Encode(books)
	}
}

func (c BookController) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Get one book")
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		logFatal(err)

		bookRepo := bookRepository.BookRepository{}
		book := bookRepo.GetBook(db, id)
		json.NewEncoder(w).Encode(book)
	}
}

func (c BookController) AddBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Add one book")
		var book models.Book
		//var bookID int
		json.NewDecoder(r.Body).Decode(&book)

		bookRepo := bookRepository.BookRepository{}
		bookID := bookRepo.AddBook(db, book)
		json.NewEncoder(w).Encode(bookID)

	}
}

func (c BookController) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Update one book")
		var book models.Book
		json.NewDecoder(r.Body).Decode(&book)

		bookRepo := bookRepository.BookRepository{}
		rowsUpdated := bookRepo.UpdateBook(db, book)
		json.NewEncoder(w).Encode(rowsUpdated)
	}
}

func (c BookController) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Remove one book")
		params := mux.Vars(r)

		id, err := strconv.Atoi(params["id"])
		logFatal(err)

		bookRepo := bookRepository.BookRepository{}
		rowsDeleted := bookRepo.RemoveBook(db, id)

		json.NewEncoder(w).Encode(rowsDeleted)

	}
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
