# Connect to a Database

## 1. Get a Database Adapter

To connect to a SQL or NoSQL database, you first need to get an adapter. To download it, use `go get` with the database type you require (`postgresql`, `mysql`, `sqlite`, `ql`, `mssql`, or `mongo`):

```
go get -u upper.io/db.v3/postgresql
```

## 2. Configure a Database Connection

Set the `ConnectionURL` using your credentials:

```go
var settings = postgresql.ConnectionURL{
	Database: `booktown`,
	Host:     `demo.upper.io`,
	User:     `demouser`,
	Password: `demop4ss`,
}
```

Note that the `ConnectionURL` (which satisfies the [db.ConnectionURL][1] interface) varies from one database type to another. The connection properties required by each adapter are explained in detail [here](https://upper.io/db.v3/adapters).


## 3. Attempt to Establish a Connection

Use the `Open` function including the `db.ConnectionURL` value we defined in the previous step: 

```go
sess, err := postgresql.Open(settings)
...
```

## 4. Set the Connection to Close

Set the database connection to close automatically after completing all tasks. Use `Close` and defer:
```
defer sess.Close()
```

[1]: https://godoc.org/upper.io/db.v3#ConnectionURL
[2]: https://godoc.org/upper.io/db.v3/lib/sqlbuilder#Open
