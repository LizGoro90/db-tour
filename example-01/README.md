# Example 01

In this example you'll learn:

* How to add an adapter to your project.
* How to connect to a database.
* How to point to a collection/table.
* How to query all items from a collection/table.
* How to use the result-set syntax.

## Setting up upper-db

1. Get the adapter with `go get`:

```
go get -u upper.io/db.v3/postgresql
```

2. Configure the adapter's `ConnectionURL`.

```go
var settings = postgresql.ConnectionURL{
	Database: `booktown`,
	Host:     `demo.upper.io`,
	User:     `demouser`,
	Password: `demop4ss`,
}
```

3. Use the adapter's `Open` function to establish a connection to the DBMS.

```go
sess, err := postgresql.Open(settings)
...
defer sess.Close()
```

## Common result-oriented syntax

upper-db provides a common result-set oriented syntax that you can use to work
with different SQL and NoSQL databases using the same API.

## Related resources

* [db.ConnectionURL](https://godoc.org/upper.io/db.v3#ConnectionURL)
* [db.Database](https://godoc.org/upper.io/db.v3#Database)
* [db.Collection](https://godoc.org/upper.io/db.v3#Collection)
