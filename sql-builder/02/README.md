# JOIN queries and struct composition

Let's suppose you have a few independent Go structs `Book`, `Author` and
`Subject`, each one mapped to a table, and that you have a JOIN query that
returns a combined result with columns from all those tables.

Create a new struct and embed all those structs using the `inline` option
(`db:",inline"`) like this:

```go
type BookAuthorSubject struct {
  Book    `db:",inline"`
  Author  `db:",inline"`
  Subject `db:",inline"`
}
```

Then build a JOIN query using the query builder:


```go
q := sess.Select("b.id AS book_id", "*").
  From("books AS b").
  Join("subjects AS s").On("b.subject_id = s.id").
  Join("authors AS a").On("b.author_id = a.id").
  OrderBy("a.last_name", "b.title")
```

Note that we're creating an alias for `book_id` with `Select("b.id AS book_id",
"*")`, this is because all three embedded structs have an ambiguous `id` field.
Our final `BookAuthorSubject` struct should look like this:

```go
type BookAuthorSubject struct {
  BookID   uint `db:"book_id"`

  Book    `db:",inline"`
  Author  `db:",inline"`
  Subject `db:",inline"`
}
```

Use the `All` method on the query to dump all its results into the `books`
slice.

```go
var books []BookAuthorSubject
err := q.All(&books)
```
