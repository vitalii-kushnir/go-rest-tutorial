package bookRepository

import (
	"database/sql"
	"go-rest-tutorial/models"
	"log"
)

type BookRepository struct{}

func (b BookRepository) GetBooks(db *sql.DB) []models.Book {
	var book models.Book
	var books = []models.Book{}

	rows, err := db.Query("select * from books")
	logFatal(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		logFatal(err)
		books = append(books, book)
	}
	return books
}

func (b BookRepository) GetBook(db *sql.DB, id int) models.Book {
	var book models.Book
	rows := db.QueryRow("select * from books where id=$1", id)
	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	logFatal(err)
	return book
}

func (b BookRepository) AddBook(db *sql.DB, book models.Book) int64 {
	var bookID int64

	err := db.QueryRow("insert into books (title, author, year) values ($1, $2, $3) RETURNING id;",
		book.Title, book.Author, book.Year).Scan(&bookID)
	logFatal(err)
	return bookID
}

func (b BookRepository) UpdateBook(db *sql.DB, book models.Book) int64 {
	result, err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 RETURNING id;",
		book.Title, book.Author, book.Year, book.ID)
	logFatal(err)

	rowsUpdated, err := result.RowsAffected()
	logFatal(err)
	return rowsUpdated
}

func (b BookRepository) RemoveBook(db *sql.DB, id int) int64 {
	result, err := db.Exec("delete from books where id=$1;", id)
	logFatal(err)

	rowsDeleted, err := result.RowsAffected()
	logFatal(err)

	return rowsDeleted
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
