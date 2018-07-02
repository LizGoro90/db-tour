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

func main() {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()

	// The Collection method returns a reference to a specific database
	// structure. In this case, the structure is a table named "books".
    col := sess.Collection("books")

	// Name and Exists are among the different methods we can call. We'll
	// use them to know if a given structure is included in the database.
	log.Println("Collection:", col.Name(), "Exists?:", col.Exists())
}
