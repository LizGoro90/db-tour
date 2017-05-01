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

	// The Collection method returns a collection reference
	col := sess.Collection("books")

	// You can use any of the methods defined on
	// https://godoc.org/upper.io/db.v3#Collection on this
	// collection reference.
	log.Println("Collection:", col.Name(), "Exists?:", col.Exists())
}
