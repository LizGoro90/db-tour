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

type Book struct {
	ID        uint   `db:"id"`
	Title     string `db:"title"`
	AuthorID  uint   `db:"author_id"`
	SubjectID uint   `db:"subject_id"`
	// Remember that the `db:"-"` tag is used for exported fields we don't want
	// to map to a column. 
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

	// Next is used to go over all items one by one. It proves useful when 
	// copying large data sets into a slice is impractical. 
	var book Book
	for res.Next(&book) {
		log.Printf("Book %d:\t%#v", book.ID, book)
	}

	// In the event of a problem, Next will break the loop and generate an
	// error (which can be retrieved by calling Err). On the other hand, when
	// the loop is successfully completed (even if the data set had no items),
	// Err will be nil.
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
