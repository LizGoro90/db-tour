# UPDATE, INSERT and DELETE queries

The `SelectFrom` method creates and returns a [Selector][1] that can be used to
build a SELECT query.

```go
var eaPoe Author

q := sess.SelectFrom("authors").
  Where("last_name", "Poe").
  One(&eaPoe)

err = q.One(&eaPoe)
```

The `Update` method creates and returns an [Updater][2] that can be used to
build an UPDATE query:

```go
q := sess.Update("authors").
  Set("first_name = ?", "Edgar Allan").
  Where("id = ?", eaPoe.ID)

res, err := q.Exec()
```

The `InsertInto` method creates and returns an [Inserter][3] that can be used
to build an INSERT query:

```go
res, err = sess.InsertInto("books").
  Columns(
    "title",
    "author_id",
    "subject_id",
  ).
  Values(
    "Brave New World",
    45,
    11,
  ).
  Exec()
```

In the example above, using `Columns` is not mandatory, the `Values` method can
also take an struct and map column-values by itself, like this:

```go
book := Book{
  Title:    "The Crow",
  AuthorID: eaPoe.ID,
}

res, err = sess.InsertInto("books").
  Values(book).
  Exec()
```

The `DeleteFrom` method creates and returns a [Deleter][4] that can be used to
build a DELETE query:

```go
q := sess.DeleteFrom("books").
  Where("title", "The Crow")

res, err := q.Exec()
```

See the
[sqlbuilder.SQLBuilder](https://godoc.org/upper.io/db.v3/lib/sqlbuilder#SQLBuilder)
interface to learn about all available methods that can be used to build or
execute SQL statements.

[1]: https://godoc.org/upper.io/db.v3/lib/sqlbuilder#Selector
[2]: https://godoc.org/upper.io/db.v3/lib/sqlbuilder#Updater
[3]: https://godoc.org/upper.io/db.v3/lib/sqlbuilder#Inserter
[4]: https://godoc.org/upper.io/db.v3/lib/sqlbuilder#Deleter

