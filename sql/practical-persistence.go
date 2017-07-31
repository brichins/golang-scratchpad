// Taken from http://www.alexedwards.net/blog/practical-persistence-sql
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Book struct {
	isbn   string
	title  string
	author string
	price  float32
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres@localhost?sslmode=disable")
	if err != nil {
		log.Fatal("Could not open database: ", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		log.Fatal("Could not query books", err)
	}
	defer rows.Close()

	bks := make([]*Book, 0)
	for rows.Next() {
		bk := new(Book)
		err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price)
		if err != nil {
			log.Fatal("Could not read row for book")
		}
		bks = append(bks, bk)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, bk := range bks {
		fmt.Printf("%s, %s, %s, $%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
	}

}
