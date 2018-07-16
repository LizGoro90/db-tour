package main

import (
	"log"

	"upper.io/db.v3"
	"upper.io/db.v3/postgresql"
)

var settings = postgresql.ConnectionURL{
	Database: `booktown`,
	Host:     `demo.upper.io`,
	User:     `demouser`,
	Password: `demopass`,
}

// Book represents an item from the "books" table. This
// table has an integer primary key ("id"):
//
// booktown=> \d books
//        Table "public.books"
//    Column   |  Type   | Modifiers
// ------------+---------+-----------
//  id         | integer | not null
//  title      | varchar | not null
//  author_id  | integer |
//  subject_id | integer |
// Indexes:
//     "books_id_pkey" PRIMARY KEY, btree (id)
//     "books_title_idx" btree (title)
type Book struct {
	ID        uint   `db:"id"`
	Title     string `db:"title"`
	AuthorID  uint   `db:"author_id"`
	SubjectID uint   `db:"subject_id"`
}

func main() {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Printf("Could not establish a connection with database: %v.", err)
		log.Fatalf(`SUGGESTION: Set password to "demop4ss" and try again.`)
	}
	defer sess.Close()

	sess.SetLogging(true)

	log.Printf("Connected to %q using %q", sess.Name(), sess.ConnectionURL())

	booksTable := sess.Collection("books")

	var book Book

	// If this table has an integer primary key you can pass
	// an int to Find and Find will look for the element that
	// matches that primary key.
	err = booksTable.Find(1).One(&book)
	if err != nil {
		if err == db.ErrNoMoreRows {
			log.Printf("One: %v", err)
			log.Printf("SUGGESTION: Change Find(1) into Find(4267).")
		} else {
			log.Printf("One: %v", err)
		}
		log.Fatal("An error ocurred, cannot continue.")
	}

	log.Printf("Book: %#v", book)
}
