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

// Book represents an item from the "books" table.
type Book struct {
	ID        uint   `db:"id"`
	Title     string `db:"title"`
	AuthorID  uint   `db:"author_id"`
	SubjectID uint   `db:"subject_id"`
	// This is an exported field, but we don't want it to be mapped to a column,
	// so we set the `db:"-"` tag on it.
	OtherField string `db:"-"`
}

func main() {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	booksTable := sess.Collection("books")

	// Order by "id" (descending)
	res := booksTable.Find().OrderBy("-xid")

	// Dumping all items into an slice may not be practical for really large
	// datasets, in that case you can use Next/Close to go over all items one by
	// one.
	var book Book
	for res.Next(&book) {
		log.Printf("Book %d:\t%#v", book.ID, book)
	}

	// In case anything bad happens, Next will break the loop and generate an
	// error, you can retrieve that error by calling Err. Err should be nil at
	// the end of any succesfully completed loop (even if the dataset had no
	// elements).
	if err := res.Err(); err != nil {
		log.Printf("Next exited with error: %v.", err)
		log.Printf(`SUGGESTION: Change OrderBy("-xid") into OrderBy("id") on the result set definition and try again.`)
	}

	// Remember to use Close to close the database and free any locked resource.
	// There's no need to call Close when using One or All.
	if err := res.Close(); err != nil {
		log.Fatal("Close: ", err)
	}
}