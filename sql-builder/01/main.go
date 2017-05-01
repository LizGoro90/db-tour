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
	ID        uint   `db:"id,omitempty"`
	Title     string `db:"title"`
	AuthorID  uint   `db:"author_id,omitempty"`
	SubjectID uint   `db:"subject_id,omitempty"`
}

func main() {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()

	sess.SetLogging(true)

	// The Collection / Find / Result syntax was created with
	// compatibility across all supported databases in mind.
	// Sometimes that won't be enough for all your needs. If
	// that's your case, you can also use the SQL builder or
	// even raw SQL depending on how complex your query is.
	//
	// The SQLBuilder is defined in the
	// upper.io/db.v3/lib/sqlbuilder package, you can see all
	// its available methods on
	// https://godoc.org/upper.io/db.v3/lib/sqlbuilder#SQLBuilder.
	q := sess.SelectFrom("books")

	// q is a sqlbuilder.Selector, you can chain any of its
	// other methods that return the same type. See
	// https://godoc.org/upper.io/db.v3/lib/sqlbuilder#Selector
	q = q.OrderBy("title")

	// Remember that queries are immutable, here p is a
	// completely independent query.
	p := q.Where("title LIKE ?", "P%")

	// Queries are not compiled nor executed until you use
	// methods like One or All.
	var booksQ, booksP []Book
	if err := q.All(&booksQ); err != nil {
		log.Fatal("q.All: ", err)
	}

	// The Iterator method is a way to iterator over large
	// sets of results one by one.
	booksP = make([]Book, 0, len(booksQ))
	iter := p.Iterator()
	var book Book
	for iter.Next(&book) {
		booksP = append(booksP, book)
	}
	// Remember to check for error values at the end of the
	// loop.
	if err := iter.Err(); err != nil {
		log.Fatal("iter.Err: ", err)
	}
	// An iterator must be closed to free up related
	// resources.
	if err := iter.Close(); err != nil {
		log.Fatal("iter.Close: ", err)
	}

	log.Printf("All books:")
	for _, book := range booksQ {
		log.Printf("Book %d:\t:%s", book.ID, book.Title)
	}
	log.Println("")

	log.Printf("Books that begin with P:")
	for _, book := range booksP {
		log.Printf("Book %d:\t:%s", book.ID, book.Title)
	}
	log.Println("")
}
