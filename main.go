package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"github.com/subosito/gotenv"

	"go-rest-tutorial/controllers"
	"go-rest-tutorial/driver"
)

func init() {
	gotenv.Load()
}

func main() {

	db := driver.ConnectDB()
	controller := controllers.BookController{}
	router := mux.NewRouter()

	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", controller.AddBooks(db)).Methods("POST")
	router.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.RemoveBook(db)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
