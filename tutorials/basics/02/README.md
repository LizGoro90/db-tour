# List All Tables/Collections in a Database

Once `Open` returns a `sess`variable (which satisfies the [db.Database][2] interface), use the `Collections` method to get all the structures in the database: 

```go
collections, err := sess.Collections()
...

log.Printf("Collections in %q: %v", sess.Name(), collections)
```

[1]: https://godoc.org/upper.io/db.v3#ConnectionURL
[2]: https://godoc.org/upper.io/db.v3#Database
[3]: https://godoc.org/upper.io/db.v3/lib/sqlbuilder#Database
[4]: https://godoc.org/upper.io/db.v3/lib/sqlbuilder#Open
