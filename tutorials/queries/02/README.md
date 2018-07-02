# Build Common Queries

Up to this point, the `Collection` method has taken the name of a structure in the database and returned a value that satisfies the [db.Collection][1] interface:

```go
booksTable := sess.Collection("books")
```

Now we'll use `Find` to search for specific objects within the hierarchy. The object
returned will be a [db.Result][2] (which begins with a condition and can contain 
zero, one, or many items.)


The `db.Result` API works the same on all supported databases, this is known as
Common Result-Oriented Syntax or CROS.

CROS is pretty useful when you want to operate on a specific set of items on a
collection or table. For instace, this is how a CROS query that fetches and
maps all the books in the "books" table looks:

```go
var books []Book

res := booksTable.Find()
err := res.All(&books)
```

This is the same query sorted by title (descending order):

```go
var books []Book

res := booksTable.Find()
res = res.OrderBy("-title")

err := res.All(&books)
```

If you only want one element from the set, use `One` instead of `All`:

```go
import "upper.io/db.v3"
...

var book Book

res := booksTable.Find(db.Cond{"id": 4})
err := res.One(&book)
```

To count all elements on the result-set use `Count`:


```go
res := booksTable.Find()

total, err := res.Count()
```

This is not the only way to query for data, depending on your database you may
have other querying APIs, for instance, SQL databases also provide a query
builder for when you need to have more control over the details of your query:

```go
q, err := sess.Select().From("books")

var books []Book
err := q.All(books)
```

And if you need absolute control over your query you can always use raw SQL:

```
rows, err := sess.Query("SELECT * FROM books")
// rows is a regular *sql.Rows object.
```

[1]: https://godoc.org/upper.io/db.v3#Collection
[2]: https://godoc.org/upper.io/db.v3#Result
