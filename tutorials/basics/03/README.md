# Create a Table/Collection Reference

Use `Collection` (also defined by [db.Database][1]) to get a database
structure in particular:

```go
col := sess.Collection("books")
```

A collection reference satisfies [db.Collection][2].

There are different methods you can call on the reference, like `Name` and `Exists`, 
which are handy for knowing whether a collection exists or not:

```go
log.Println("Collection:", col.Name(), "Exists?:", col.Exists())
```

![Note](https://github.com/LizGoro90/db-tour/tree/master/static/img)
< You can create references for collections that do not
< necessarily exist yet.

[1]: https://godoc.org/upper.io/db.v3#Database
[2]: https://godoc.org/upper.io/db.v3#Collection
