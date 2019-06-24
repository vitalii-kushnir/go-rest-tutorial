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
		books, err := bookRepo.GetBooks(db)

		if err != nil {
			handleError(w, err, http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(books)
	}
}

func (c BookController) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Get one book")
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])

		if err != nil {
			handleError(w, err, http.StatusBadRequest)
			return
		}

		bookRepo := bookRepository.BookRepository{}
		book, err := bookRepo.GetBook(db, id)
		if err != nil {
			handleError(w, err, http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(book)
	}
}

func (c BookController) AddBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Add one book")
		var book models.Book
		json.NewDecoder(r.Body).Decode(&book)

		bookRepo := bookRepository.BookRepository{}
		bookID, err := bookRepo.AddBook(db, book)

		if err != nil {
			handleError(w, err, http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(bookID)

	}
}

func (c BookController) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Update one book")
		var book models.Book
		json.NewDecoder(r.Body).Decode(&book)

		bookRepo := bookRepository.BookRepository{}
		rowsUpdated, err := bookRepo.UpdateBook(db, book)

		if err != nil {
			handleError(w, err, http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(rowsUpdated)
	}
}

func (c BookController) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Remove one book")
		params := mux.Vars(r)

		id, err := strconv.Atoi(params["id"])
		if err != nil {
			handleError(w, err, http.StatusBadRequest)
			return
		}

		bookRepo := bookRepository.BookRepository{}
		rowsDeleted, err := bookRepo.RemoveBook(db, id)
		if err != nil {
			handleError(w, err, http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(rowsDeleted)

	}
}

func handleError(w http.ResponseWriter, err error, statusCode int) {
	var errorDto models.Error
	errorDto.Message = err.Error()
	errorDto.StatusCode = statusCode

	json.NewEncoder(w).Encode(errorDto)
}
