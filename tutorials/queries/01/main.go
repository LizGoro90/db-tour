package main

import (
	"log"

	"upper.io/db.v3/postgresql"
)

var settings = postgresql.ConnectionURL{
	Database: `booktown`,
	Host:     `demo.upper.io`,
	User:     `demouser`,
	Password: `demop4ss`,
}

// The struct represents an item of the "books" table. The fields 
// accompanying the item represent the columns in the table and are
// mapped to Go values below.
type Book struct {
	// upper-db will convert Go types into database-specific types
	// and vice versa, as shown in the "books" table on the left.
	ID uint `db:"id"`
	Title string `db:"title"`
	AuthorID uint `db:"author_id"`
	SubjectID uint `db:"subject_id"`
}

func main() {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	booksTable := sess.Collection("books")

}
