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
}

func main() {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()


	// We've pointed to a database structure (a table, in this
	// case) through a collection reference.
	nonexistentTable := sess.Collection("nonexistent_table")

	// For the following method, make sure your DBMS allows to
	// use collections or tables that may not exist.
	_, err = nonexistentTable.Insert(&Book{})
	if err != nil {
		log.Println("nonexistent_table: ", err)
	}

	// Use Exists if you must check for table existence.
	if !nonexistentTable.Exists() {
		log.Println("The nonexistent_table does not exist.")
	}

	// "Books" is a table that already exists in our test
	// database.
	booksTable := sess.Collection("books")

	// We use Find to create a result set (db.Result).
	res := booksTable.Find()

	// The result set can be modified by chaining different
	// db.Result methods (like Where, And, OrderBy, Select
	// Limit, and Group). These methods will return a new
	// result set.
	res = res.OrderBy("-title") // ORDER BY title DESC

	// The result set is lazy, meaning that the query will
	// be built or sent to the database until one of the 
	// methods that require database interaction is used (for
	// example, One or All).
	var books []Book
	if err := res.All(&books); err != nil {
		log.Fatal(err)
	}

	// The All method copies every single item in the result
	// set into a Go slice.
	log.Printf("Items in %q table:\n", booksTable.Name())
	for _, book := range books {
		log.Printf("Item %d:\t%q\n", book.ID, book.Title)
	}

	// The result set can be reused many times. We're now going 
	// to count all the items in the result set.
	total, err := res.Count()
	if err != nil {
		log.Fatal("Count: ", err)
	}
	log.Printf("There are %d items on %q", total, booksTable.Name())

	// We can also build new result sets from old ones. 
	itemsThatBeginWithP := res.And("title LIKE", "P%") // WHERE ... AND title LIKE 'P%'

	// The old result set is not altered,
	total1, err := res.Count()
	if err != nil {
		log.Fatal("Count: ", err)
	}

	// The new result set is the one that was modified.
	total2, err := itemsThatBeginWithP.Count()
	if err != nil {
		log.Fatal("Count: ", err)
	}

	log.Printf("There are still %d items on %q", total1, booksTable.Name())
	log.Printf("And there are %d items on %q that begin with \"P\"", total2, booksTable.Name())
}
