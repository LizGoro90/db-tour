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

	// Use Find on a collection to create a db.Result result
	// set. See https://godoc.org/upper.io/db.v3#Result for
	// all methods on result sets.
	res := booksTable.Find()

	// A result set can be modified by chaining any of the
	// other db.Result methods that return a new db.Result,
	// like Where, And, OrderBy, Select, Limit and Group.
	res = res.OrderBy("-title") // ORDER BY title DESC

	// A result set is lazy and does not build not send query
	// to the database until you use one of the methods that
	// require interaction with the database. Like One or All.
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

	// A result set can be reused many times, here we're
	// counting all the items in the result set.
	total, err := res.Count()
	if err != nil {
		log.Fatal("Count: ", err)
	}
	log.Printf("There are %d items on %q", total, booksTable.Name())

	// We can build new result sets upon old ones.
	itemsThatBeginWithP := res.And("title LIKE", "P%") // WHERE ... AND title LIKE 'P%'

	// The original result set is not affected.
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
