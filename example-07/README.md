# How to embed structs and create a JOIN query

Let's suppose you wrote three Go structs to represent three different tables:
`Book`, `Author` and `Subject`; and that you want to create a join query that
returns a result with columns from all these tables.

You can create a new struct and embed these three parent structs by using the
`inline` option (`db:",inline"`) like this:

```go
type BookAuthorSubject struct {
  Book    `db:",inline"`
  Author  `db:",inline"`
  Subject `db:",inline"`
}
```

Then you can use the query builder to create a query with a join between these
three tables:

```go
q := sess.Select("b.id AS book_id", "*").
  From("books AS b").
  Join("subjects AS s").On("b.subject_id = s.id").
  Join("authors AS a").On("b.author_id = a.id").
  OrderBy("a.last_name", "b.title")
```

Note that we're setting an alias for `book_id`, this is because all three
embedded structs have an `id` field. Our final `BookAuthorSubject` should look
like this:

```go
type BookAuthorSubject struct {
  BookID   uint `db:"book_id"`

  Book    `db:",inline"`
  Author  `db:",inline"`
  Subject `db:",inline"`
}
```

Finally, use the `All` method on the query to dump all results into the `books`
variable.

```go
var books []BookAuthorSubject

err := q.All(&books)
```
