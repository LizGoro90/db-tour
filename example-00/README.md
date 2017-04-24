# How to connect to a database and list all its tables

In this example you'll learn:

* How to get an upper-db adapter.
* How to add an adapter to your project.
* How to connect to a database.
* How to get a list of all tables within the database.

## Get an adapter

Get a database adapter with `go get`:

```
go get -u upper.io/db.v3/postgresql
```

Besides `postgresql`, you can try with `mysql`, `sqlite`, `ql`, `mssql` and
`mongo`.


## Configure a database connection

Configure the adapter's `ConnectionURL`.

```go
var settings = postgresql.ConnectionURL{
	Database: `booktown`,
	Host:     `demo.upper.io`,
	User:     `demouser`,
	Password: `demop4ss`,
}
```

Not all databases require the same information to connect to or open a
database, this is why all adapters provide their own `ConnectionURL` type
(which satisfies the [db.ConnectionURL][1] interface). You can see a
description of your adapter's particular `ConnectionURL` at godoc.org:
https://godoc.org/upper.io/db.v3/postgresql#ConnectionURL

## Use the `Open` function to attempt to establish a connection

Use the adapter's `Open` function to attempt to establish a connection with the
database.

All adapters have an [Open][4] function that accepts a `db.ConnectionURL`, like
the adapter's own `ConnectionURL` we created on the previous step:

```go
// Attempt to establish a connection with a database.
sess, err := postgresql.Open(settings)
...

// Remember to close the database when you're no longer using it.
defer sess.Close()
```

## Using [db.Database][2] methods

The `sess` variable returned by `Open` is a database reference that satisfies
[db.Database][2]. If you're working with a SQL database, `sess` will also
satisfy [sqlbuilder.Database][3].

The [db.Database][2] interface defines a `Collection` method. Use the
`Collection` method to get all the names of the collections in the database:

```go
collections, err := sess.Collections()
if err != nil {
	log.Fatal("Collections: ", err)
}
```

## How to get all items on a collection

Go to our [next example](/tour/01) to see how to get and map all items on a
collection and map them into a Go type.

[1]: https://godoc.org/upper.io/db.v3#ConnectionURL
[2]: https://godoc.org/upper.io/db.v3#Database
[3]: https://godoc.org/upper.io/db.v3/lib/sqlbuilder#Database
[4]: https://godoc.org/upper.io/db.v3/lib/sqlbuilder#Open
