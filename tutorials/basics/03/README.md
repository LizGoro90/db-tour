# Create a table/collection reference

`Collection` is one of the methods defined by [db.Database][1]. This
method can be used to create a collection reference.

```go
col := sess.Collection("books")
```

A collection reference satisfies [db.Collection][2].

`Name` and `Exists` are two methods defined by `db.Collection`.

```go
log.Println("Collection:", col.Name(), "Exists?:", col.Exists())
```

Keep in mind that you can create references for collections that do not
necessarily exist yet.

[1]: https://godoc.org/upper.io/db.v3#Database
[2]: https://godoc.org/upper.io/db.v3#Collection
