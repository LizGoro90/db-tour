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

// Author represents an item from the "authors" table.
type Author struct {
	ID        uint   `db:"id,omitempty"`
	LastName  string `db:"last_name"`
	FirstName string `db:"first_name"`
}

// Subject represents an item from the "subjects" table.
type Subject struct {
	ID       uint   `db:"id,omitempty"`
	Subject  string `db:"subject"`
	Location string `db:"location"`
}

func main() {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()

	sess.SetLogging(true)

	var eaPoe Author

	// Using sqlbuilder.Selector to get E .A. Poe from our
	// authors table.
	err = sess.SelectFrom("authors").
		Where("last_name", "Poe"). // Or Where("last_name = ?", "Poe")
		One(&eaPoe)
	if err != nil {
		log.Fatal("Query: ", err)
	}
	log.Printf("%#v", eaPoe)

	// The name says "Edgar Allen", let's fit it using
	// sqlbuilder.Updater:
	res, err := sess.Update("authors").
		Set("first_name = ?", "Edgar Allan"). // Or Set("first_name", "Edgar Allan").
		Where("id = ?", eaPoe.ID).            // Or Where("id", eaPoe.ID)
		Exec()
	if err != nil {
		log.Printf("Query: %v. This is expected on the read-only sandbox", err)
	}

	// Now let's create a new E. A. P. book.
	book := Book{
		Title:    "The Crow",
		AuthorID: eaPoe.ID,
	}
	res, err = sess.InsertInto("books").
		Values(book). // Or Columns(c1, c2, c2, ...).Values(v1, v2, v2, ...).
		Exec()
	if err != nil {
		log.Printf("Query: %v. This is expected on the read-only sandbox", err)
	}
	if res != nil {
		id, _ := res.LastInsertId()
		log.Printf("New book id: %d", id)
	}

	// Delete the book we just created (and any book with the
	// same name).
	q := sess.DeleteFrom("books").
		Where("title", "The Crow")
	log.Printf("Compiled query: %v", q)

	_, err = q.Exec()
	if err != nil {
		log.Printf("Query: %v. This is expected on the read-only sandbox", err)
	}
}