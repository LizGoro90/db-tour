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

	// The settings variable has a String method that builds
	// and returns a valid DSN. This DSN may be different
	// depending on the database you're connecting to.
	log.Print("Connected to PostgreSQL using the following DSN: ", settings.String())

	// The Collections method returns all the structures in
	// the database. In this case, the structures are tables.
	collections, err := sess.Collections()
	if err != nil {
		log.Fatal("Collections: ", err)
	}

	// Name is one of the different methods we can call. 
	// We'll use it to retrieve the database name.
	log.Printf("Collections in %q: %v", sess.Name(), collections)
}
