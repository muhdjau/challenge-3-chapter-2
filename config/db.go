package config

import (
	"challenge-chapter-2-sesi-3/models"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "365442"
	dbname   = "db_book"
)

var (
	db  *sql.DB
	err error
)

func ConnectDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success connect to database!")
}

func CreateBook(book models.Books) (newBook models.Books, err error) {
	var lastID int

	getLastId := `
		SELECT id FROM books 
		ORDER BY id DESC LIMIT 1
	`
	err = db.QueryRow(getLastId).Scan(&lastID)
	if err != nil {
		book.ID = 1
	}

	book.ID = lastID + 1

	sqlStatement := `
		INSERT INTO books (id, title, author, description) 
		VALUES ($1, $2, $3, $4) 
		Returning *
	`
	_, err = db.Exec(sqlStatement, book.ID, book.Title, book.Author, book.Description)
	if err != nil {
		return
	}

	newBook = book
	return
}

func GetAllBooks(allBook []models.Books) (allBooks []models.Books, err error) {
	sqlStatement := `SELECT * FROM books ORDER BY id`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var books models.Books

		err = rows.Scan(&books.ID, &books.Title, &books.Author, &books.Description)

		if err != nil {
			log.Println("Error di db query")
			return
		}

		allBooks = append(allBooks, books)
	}

	return
}

func GetBookById(bookID int, book models.Books) (bookData models.Books, err error) {
	sqlStatement := `
		SELECT * FROM books 
		WHERE id = $1
	`
	err = db.QueryRow(sqlStatement, bookID).Scan(&bookData.ID, &bookData.Title, &bookData.Author, &bookData.Description)
	if err != nil {
		return
	}

	return
}

func UpdateBook(bookID int, updateBook models.Books) (book models.Books, err error) {
	findId := `
		SELECT id FROM books 
		WHERE id = $1
	`

	err = db.QueryRow(findId, bookID).Scan(&bookID)
	if err != nil {
		return
	}

	sqlStatement := `
		UPDATE books 
		SET title = $2, author = $3, description = $4 
		WHERE id = $1
	`
	_, err = db.Exec(sqlStatement, bookID, updateBook.Title, updateBook.Author, updateBook.Description)
	if err != nil {
		return
	}

	book = updateBook
	book.ID = bookID
	return
}

func DeleteBook(idBook int) (err error) {
	findId := `
		SELECT id FROM books 
		WHERE id = $1
	`
	err = db.QueryRow(findId, idBook).Scan(&idBook)
	if err != nil {
		return
	}

	sqlStatement := `
		DELETE FROM books 
		WHERE id = $1
	`

	_, err = db.Exec(sqlStatement, idBook)
	if err != nil {
		panic(err)
	}

	return
}
