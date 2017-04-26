# How to list all tables/collections in a database

The `sess` variable returned by `Open` satisfies the [db.Database][2]
interface. If you're working with a SQL database, `sess` will satisfy
[sqlbuilder.Database][3] as well.

Use the `Collection` method (defined by [db.Database][2]) to get the names of
all the collections in the database:

```go
collections, err := sess.Collections()
...

log.Printf("Collections in %q: %v", sess.Name(), collections)
```

## How to get all items on a collection

Go to our [next example](/tour/01) to see how to get and map all items on a
collection and map them into a Go type.

[1]: https://godoc.org/upper.io/db.v3#ConnectionURL
[2]: https://godoc.org/upper.io/db.v3#Database
[3]: https://godoc.org/upper.io/db.v3/lib/sqlbuilder#Database
[4]: https://godoc.org/upper.io/db.v3/lib/sqlbuilder#Open
