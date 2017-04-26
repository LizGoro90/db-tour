package main

import (
	"log"

	// In this example we're going to connect to a PostgreSQL database, for this
	// we need to import the postgresql adapter. If you want to try with another
	// DBMS, change the adapter name and provide valid connection settings.
	"upper.io/db.v3/postgresql"
)

// The postgresql adapter provides a ConnectionURL type, this type has
// everything you need to establish a connection to a PostgreSQL database. The
// postgresql.ConnectionURL type satisfies db.ConnectionURL
// (https://godoc.org/upper.io/db.v3#ConnectionURL).
var settings = postgresql.ConnectionURL{
	Database: `booktown`,
	Host:     `demo.upper.io`,
	User:     `demouser`,
	Password: `demop4ss`,
}

func main() {
	// Use the Open function of the adapter to request a connection to the
	// database. Open returns a sqlbuilder.Database type (all SQL adapters do the
	// same), refer to https://godoc.org/upper.io/db.v3/lib/sqlbuilder#Database
	// for all available methods.
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()

	log.Printf("Connected!")
}
