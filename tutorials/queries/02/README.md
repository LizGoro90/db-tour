# Build Common Queries

Up to this point, the `Collection` method has taken the name of a structure in the database and returned a value that satisfies the [db.Collection][1] interface:

```go
booksTable := sess.Collection("books")
```

Now we'll use `Find` to search for specific objects within the hierarchy. The object
returned will be a [db.Result][2] (which begins with a condition and can contain 
zero, one, or many items.)

![Note](https://github.com/LizGoro90/db-tour/tree/master/static/img)
< The `db.Result` API works the same on all supported databases, this is known as
< Common Result-Oriented Syntax (CROS) and comes in handy when you want to query
< a collection or a table.

For instance, the following is a CROS query that fetches and
maps all the items in the "books" table:

```go
var books []Book

res := booksTable.Find()
err := res.All(&books)
```

You can build the query to return items in different ways, such as sorted by title (descending order):

```go
var books []Book

res := booksTable.Find()
res = res.OrderBy("-title")

err := res.All(&books)
```

... or use `One` instead of `All` if you want to retrieve a single item from the set:

```go
import "upper.io/db.v3"
...

var book Book

res := booksTable.Find(db.Cond{"id": 4})
err := res.One(&book)
```

You can also determine the total number of items in the result-set with `Count`:


```go
res := booksTable.Find()

total, err := res.Count()
```

There are many options for you to define queries depending on your database type. Take
a look [here](https://upper.io/db.v3/getting-started#defining-a-result-set-with-code-find-code).

![Note](https://github.com/LizGoro90/db-tour/tree/master/static/img)
> In the particular case of SQL databases, you can also choose to use a query builder 
> (for more control over your query):
```go
q, err := sess.Select().From("books")

var books []Book
err := q.All(books)
```

> ... or raw SQL (for absolute control over your query):

```
rows, err := sess.Query("SELECT * FROM books")
// rows is a regular *sql.Rows object.
```

Given that the example in this tour is based on a SQL database, we'll elaborate hereunder on the use of both a) the SQL builder and b) raw SQL.

![Note](https://github.com/LizGoro90/db-tour/tree/master/static/img)
> If you're working with a NoSQL database, refer to next page to continue learning about 
> querying.

[1]: https://godoc.org/upper.io/db.v3#Collection
[2]: https://godoc.org/upper.io/db.v3#Result
