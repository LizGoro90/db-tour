package main

import (
	"log"

	// In this example, we're going to connect to a PostgreSQL
	// database.
	"upper.io/db.v3/postgresql"
)

// We set the ConnectionURL type required by PostgreSQL using
// the reference documentation.
var settings = postgresql.ConnectionURL{
	Database: `booktown`,
	Host:     `demo.upper.io`,
	User:     `demouser`,
	Password: `demop4ss`,
}

func main() {
	// Use the Open function of the adapter to request a
	// connection to the database.  Open returns a
	// sqlbuilder.Database type (all SQL adapters do the
	// same), refer to
	// https://godoc.org/upper.io/db.v3/lib/sqlbuilder#Database
	// for all available methods.
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()

	log.Printf("Connected!")
}
