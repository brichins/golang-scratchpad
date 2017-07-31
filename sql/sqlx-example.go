package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	ISBN   string
	Title  string
	Author string
	Price  float32
}

var db *sqlx.DB

func main() {
	db = sqlx.MustConnect("sqlite3", ":memory:")
	sql1 := `CREATE TABLE books (
  isbn    char(14) NOT NULL PRIMARY KEY,
  title   varchar(255) NOT NULL,
  author  varchar(255) NOT NULL,
  price   decimal(5,2) NOT NULL
);`
	result := db.MustExec(sql1)
	size, err := result.RowsAffected()
	log.Printf("Affected %v rows  Err: %v\n", size, err)

	sql2 := `INSERT INTO books (isbn, title, author, price) VALUES
('978-1503261969', 'Emma', 'Jayne Austen', 9.44),
('978-1505255607', 'The Time Machine', 'H. G. Wells', 5.99),
('978-1503379640', 'The Prince', 'Niccolò Machiavelli', 6.99);`
	result = db.MustExec(sql2)
	size, err = result.RowsAffected()
	log.Printf("Affected %v rows  Err: %v\n", size, err)

	books := []Book{}
	err = db.Select(&books, "SELECT * FROM books")
	if err != nil {
		log.Fatal("Error selecting books || ", err)
	}

	log.Printf("Books: %+v\n", books)
}
