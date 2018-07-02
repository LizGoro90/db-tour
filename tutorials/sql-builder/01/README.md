## a) SQL Builder

All the adapters for SQL databases come with a handy builder that you
can use to compose queries, as defined [here](https://upper.io/db.v3/lib/sqlbuilder).

There are different [SQLBuilder][1] methods you can use to begin your query, for
instance `SelectFrom`:

```go
q := sess.SelectFrom("books")
```

`SelectFrom` returns a [Selector][2] and some `Selector` methods return `Selector` 
too, so you can chain method calls like this:

```go
q := sess.SelectFrom("books").
  Where("title LIKE ?", "P%")
```

or

```go
q := sess.SelectFrom("books")
q = q.Where("title LIKE ?", "P%")
```

Note that in the example above where reassigning `q`, this is because queries
are immutable and methods do not affect the caller. For instance, in the
following example `q` is not affected by `Where`:

```go
q := sess.SelectFrom("books")
p := q.Where("title LIKE ?", "P%").
  OrderBy("title")
```

Use the `All` or `One` methods on a query to compile, execute and map its
results into a Go type:

```go
var books []Book
err := q.All(&books)
```

Or

```go
var book Book
err := q.All(&book)
```

The `Selector` interface also provides an special `Iterator` method that you
can use to create an iterator and query over results one by one:

```go
iter := q.Iterator()

for iter.Next(&book) {
  // ...
}

if err := iter.Err(); err != nil {
  // ...
}

if err := iter.Close(); err != nil {
  // ...
}
```

Once you're done with the iterator remember to check for any errors with `Err`
and to free any locked resources with `Close`.

[1]: https://godoc.org/upper.io/db.v3/lib/sqlbuilder#SQLBuilder
[2]: https://godoc.org/upper.io/db.v3/lib/sqlbuilder#Selector
