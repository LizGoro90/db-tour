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

// Book represents an item from the "books" table, column
// names are mapped to Go values.
type Book struct {
	// Map the "id" column to the ID field. Only exported
	// fields can be mapped to database columns.
	ID uint `db:"id"`
	// The "title" column is a VARCHAR type, upper-db converts
	// Go types into database-specific types and vice versa.
	Title string `db:"title"`
	// The "author_id" column is an integer type.
	AuthorID uint `db:"author_id"`
	// The "subject_id" column is an integer type.
	SubjectID uint `db:"subject_id"`
}

func main() {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	booksTable := sess.Collection("books")

	res := booksTable.Find().OrderBy("id")

	var books []Book
	if err := res.All(&books); err != nil {
		log.Fatal(err)
	}

	// The All method dumps all the items in the result set
	// into a Go slice.
	log.Printf("Items in %q table:\n", booksTable.Name())
	for _, book := range books {
		log.Printf("Item %d:\t%q\n", book.ID, book.Title)
	}
}
