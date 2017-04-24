# How to query and map large datasets

In this chapter you'll learn:

* How to exclude struct fields from a query.
* How to query results one by one.

## Exclude a struct field

If you want to exclude an exported field from a query, set the `db` field tag
to `-`:

```go
type Book struct {
	OtherField string `db:"-"`
}
```

## Map results one by one

Suppose that you have a very large dataset and that mapping all matching items
into an slice is impractical for memory and performance reasons.

The `Next` method allows items on a `db.Result` to be mapped one by one:

```go
res := booksTable.Find().OrderBy("-id")

var book Book
for res.Next(&book) {
  // ...
}
```

When a result set does not have more items `Next` will return false and break
the loop.

When dealing with results one by one you'll also need to check for errors and
free locked resources manually, unlike `One` or `All`.

```go
if err := res.Err(); err != nil {
  ...
}

if err := res.Close(); err != nil {
  ...
}
```
