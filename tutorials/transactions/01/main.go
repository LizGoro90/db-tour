package main

import (
	"context"
	"log"

	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
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

	// The ctx value will be passed to sess.Tx. You can define
	// any context.Context or pass nil instead. 
	ctx := context.Background()

	// sess.Tx requires a function, this function takes a
	// single sqlbuilder.Tx argument and returns an error. The
	// tx value is just like sess, except it lives within a
	// transaction.  If the function returns any error, the
	// transaction will be rolled back.
	err = sess.Tx(ctx, func(tx sqlbuilder.Tx) error {
		// Anything you do here with the tx value will be part
		// of the transaction.
		cols, err := tx.Collections()
		if err != nil {
			return err
		}
		log.Printf("Cols: %#v", cols)

		// The booksTable value is valid only within the transaction.
		booksTable := tx.Collection("books")
		total, err := booksTable.Find().Count()
		if err != nil {
			return err
		}
		log.Printf("There are %d items in %s", total, booksTable.Name())

		var books []Book
		err = tx.SelectFrom("books").Limit(3).OrderBy(db.Raw("RANDOM()")).All(&books)
		if err != nil {
			return err
		}
		log.Printf("Books: %#v", books)

		res, err := tx.Query("SELECT * FROM books ORDER BY RANDOM() LIMIT 1")
		if err != nil {
			return err
		}

		var book Book
		err = sqlbuilder.NewIterator(res).One(&book)
		if err != nil {
			return err
		}
		log.Printf("Random book: %#v", book)

		// If the function returns no error the transaction is commited.
		return nil
	})

	if err != nil {
		log.Printf("sess.Tx: ", err)
	}
}
