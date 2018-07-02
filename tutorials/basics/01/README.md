# Connect to a Database

## 1. Get a Database Adapter

To connect to a SQL or NoSQL database, you first need to get an adapter. To download it, use `go get` with the database type you require (`postgresql`, `mysql`, `sqlite`, `ql`, `mssql`, or `mongo`):

```
go get -u upper.io/db.v3/postgresql
```

## Configure a database connection

Configure the adapter's `ConnectionURL` using your credentials.

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
description of your adapter's particular `ConnectionURL` at
`https://upper.io/db.v2/[adapter]`, for instance
[https://upper.io/db.v3/postgresql](https://upper.io/db.v3/postgresql).

## Use the `Open` function to attempt to establish a connection

Use the adapter's `Open` function to attempt to establish a connection with the
database.

All adapters have an [Open][2] function that accepts a `db.ConnectionURL`
value, like the adapter's own `ConnectionURL` we created on the previous step:

```go
sess, err := postgresql.Open(settings)
...
```

Remember to close the database when you're no longer using it:
```
defer sess.Close()
```

[1]: https://godoc.org/upper.io/db.v3#ConnectionURL
[2]: https://godoc.org/upper.io/db.v3/lib/sqlbuilder#Open
