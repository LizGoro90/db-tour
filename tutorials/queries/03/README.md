# Query Large Data Sets

If you're working with significantly large data sets, copying all matching items
into a slice might be impractical for memory and performance reasons.

In this case, you can use `Next` to map the items in the `db.Result` one by one:

```go
res := booksTable.Find().OrderBy("-id")

var book Book
for res.Next(&book) {
  // ...
}
```

When a result set does not have any more items, `Next` will return `false` and
that'll break the loop.

When dealing with results one by one you'll also need to check for errors (with `Err()`) and
free locked resources manually (with `Close`).

```go
if err := res.Err(); err != nil {
  ...
}

if err := res.Close(); err != nil {
  ...
}
```
