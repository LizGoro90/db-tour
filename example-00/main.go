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

	// The settings variable has a String method that builds and returns a valid
	// DSN. This DSN may be different depending on the database you're connecting
	// to.
	log.Print("Connected to PostgreSQL using the following DSN: ", settings.String())

	// The Collections method returns all the collections (tables) on a database,
	// this method is part of the db.Database interface (which provides methods
	// that work on both SQL and NoSQL databases), since sqlbuilder.Database is
	// built upon db.Database, you can call all db.Database methods on a
	// sqlbuilder.Database too. See https://godoc.org/upper.io/db.v3#Database for
	// all available db.Database methods.
	collections, err := sess.Collections()
	if err != nil {
		log.Fatal("Collections: ", err)
	}

	// The Name method is also part of the db.Database interface and returns the
	// name of the database you're connected to.
	log.Printf("Collections in %q: %v", sess.Name(), collections)
}
