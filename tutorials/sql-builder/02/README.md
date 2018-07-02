### JOIN Queries and Struct Composition

Now let's suppose you have independent Go structs (each one mapped to a different
table) and a JOIN query that returns a result combining columns from all the 
mapped tables.

In this scenario, you can create a new struct and embed all the structs using the
`inline` option (`db:",inline"`), like in the following example (where `Book`,
`Author`, and `Subject` are independent structs):

```go
type BookAuthorSubject struct {
  Book    `db:",inline"`
  Author  `db:",inline"`
  Subject `db:",inline"`
}
```

... and then create the JOIN query using the builder:


```go
q := sess.Select("b.id AS book_id", "*").
  From("books AS b").
  Join("subjects AS s").On("b.subject_id = s.id").
  Join("authors AS a").On("b.author_id = a.id").
  OrderBy("a.last_name", "b.title")
```

Note that an alias for `book_id` is created with `Select("b.id AS book_id","*")`. 
This is because all three embedded structs have a field with the same name (`id`),
which is ambiguous. The final struct (`BookAuthorSubject`) should look like this:

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
![Note](https://github.com/LizGoro90/db-tour/tree/master/static/img)
> An alias exists only for the duration of the query. 


```go
var books []BookAuthorSubject
err := q.All(&books)
```
