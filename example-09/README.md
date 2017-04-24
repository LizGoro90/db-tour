# How to use raw SQL

The [SQLBuilder][1] interface provides `Query`, `QueryRow` and `Exec` methods
that mimic the database/sql API:

```go
rows, err := sess.Query(`SELECT id, first_name, last_name FROM authors WHERE last_name = ?`, "Poe")
...

row, err := sess.QueryRow(`SELECT * FROM authors WHERE id = ?`, 23)
...

res, err := sess.Exec(`UPDATE authors SET first_name = ? WHERE id = ?`, "Edgar Allan", eaPoe.ID)
...
```

Feel free to use raw SQL whenever you feel like you need it. Using SQL does not
mean you'll have to map Go fields by hand, you can import the
`upper.io/db.v3/lib/sqlbuilder` package and use the `sqlbuilder.NewIterator`
function:

```go
iter := sqlbuilder.NewIterator(rows)

var books []Book
err := iter.All(&books)
```

This iterator provides well-known upper-db methods like `One`, `All` and
`Next`.

[1]: https://godoc.org/upper.io/db.v3/lib/sqlbuilder#SQLBuilder
