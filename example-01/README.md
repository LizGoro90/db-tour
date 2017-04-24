# How to map Go and database types

In this example you'll learn:

* How to create a mapping between Go types and database types.
* How to create a reference to a collection/table.
* How to query all items from a collection/table.
* How to use the result-set syntax.

## Mapping Go types and database types

You can use the `db` field tag to tell upper-db which database columns
correspond to specific struct fields. The general syntax looks like:

```
type Whatever struct {
	FieldName fieldType `db:"field_name,options"`
}
```

Please keep in mind that:

* The field and column must be compatible types (upper-db will handle
	conversions automatically).
* The field must be an exported field.

For instance, let's suppose we created a table like this:

```sql
CREATE TABLE "books" (
	"id" INTEGER NOT NULL,
	"title" TEXT NOT NULL,
	"author_id" INTEGER,
	"subject_id" INTEGER,
	CONSTRAINT "books_id_pkey" PRIMARY KEY ("id")
);
```

Then we can start by creating a `Book` struct representing one element from
said table:

```go
type Book struct {

}
```

we can add all the fields we want to use:

```go
type Book struct {
	ID         uint
	Title      string
	AuthorID   uint
	SubjectID  uint
}
```

finally, we add `db` tags to the fields we want to map:

```go
type Book struct {
	ID          uint   `db:"id"`
	Title       string `db:"title"`
	AuthorID    uint   `db:"author_id"`
	SubjectID   uint   `db:"subject_id"`
}
```

## Mapping fields with auto-values

Some databases can be configured to insert automatically-generated IDs,
serials, dates or other values. If the table you're working with has a column
like that you'll need to add the `omitempty` option to its `db` tag:

```go
type Book struct {
	ID uint `db:"id,omitempty"`
}
```

The `omitempty` option will make upper-db ignore zero-valued fields on INSERTs
and UPDATEs.

The following table describes all available options for `db` tags.

| Option           | Description                                     |
| :--------------- | :---------------------------------------------- |
| `omitempty`      | The field is skipped when zero |

## Creating a table reference

The `db.Database` interface defines a `Collection` method. `Collection` takes
the name of a collection or a table name and returns a reference to it:

```go
booksTable := sess.Collection("books")
```

Collections references satisfy the [db.Collection][1] interface.

## Common result-oriented syntax (CROS)

One of the methods defined by `db.Collection` is `Find`. With `Find` you can
create a [db.Result][2] reference.

Result-sets start with a condition and they can contain zero, one or many
items. You can operate on result-sets with [db.Result][2] methods.

The `db.Result` API work the same on all supported databases, this is known as
Common Result-Oriented Syntax or CROS.

CROS is pretty useful when you're querying for items, this is how a CROS query
that fetches and maps all books in the "books" table would look:

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
var book Book

res := booksTable.Find()
err := res.One(&book)
```

To count all elements on the result-set use `Count`:


```go
res := booksTable.Find()
res = res.OrderBy("-title")

total, err := res.Count()
```

Depending on your database you may have other querying APIs, for instance, SQL
databases also provide a query builder for when you need to have more control
over the details of your query:

```go
q, err := sess.Select().From("books")

var books []Book
err := q.All(books)
```

And if you need absolute control over your results you can always use raw SQL:

```
rows, err := sess.Query("SELECT * FROM books")
// rows is a regular *sql.Rows object.
```

## Querying large datasets

Keep on reading to know [how to query and map large datasets](/tour/02).

[1]: https://godoc.org/upper.io/db.v3#Collection
[2]: https://godoc.org/upper.io/db.v3#Result
