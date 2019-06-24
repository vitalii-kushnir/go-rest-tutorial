package bookRepository

import (
	"database/sql"
	"go-rest-tutorial/models"
)

type BookRepository struct{}

func (b BookRepository) GetBooks(db *sql.DB) ([]models.Book, error) {
	var book models.Book
	var books = []models.Book{}

	rows, err := db.Query("select * from books")
	if err != nil {
		return books, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		if err != nil {
			return books, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (b BookRepository) GetBook(db *sql.DB, id int) (models.Book, error) {
	var book models.Book
	rows := db.QueryRow("select * from books where id=$1", id)
	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	if err != nil {
		return book, err
	}
	return book, nil
}

func (b BookRepository) AddBook(db *sql.DB, book models.Book) (int64, error) {
	var bookID int64

	err := db.QueryRow("insert into books (title, author, year) values ($1, $2, $3) RETURNING id;",
		book.Title, book.Author, book.Year).Scan(&bookID)
	if err != nil {
		return 0, err
	}
	return bookID, nil
}

func (b BookRepository) UpdateBook(db *sql.DB, book models.Book) (int64, error) {
	result, err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 RETURNING id;",
		book.Title, book.Author, book.Year, book.ID)
	if err != nil {
		return 0, err
	}

	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsUpdated, nil
}

func (b BookRepository) RemoveBook(db *sql.DB, id int) (int64, error) {
	result, err := db.Exec("delete from books where id=$1;", id)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsDeleted, nil
}
