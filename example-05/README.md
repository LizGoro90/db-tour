# How to insert new items into a table / collection

In order to insert a new item into a collection you'll first need a reference
to that collection:

```go
booksTable := sess.Collection("books")
```

Use the `Insert` method on a collection reference to insert an item
without changing the passed value:

```go
book := Book{
  Title:    "Kokoro",
  AuthorID: 123,
}

id, err := booksTable.Insert(book)
```

Use the `InsertReturning` method on a collection reference to insert an item
and update all its auto-fields, like IDs or dates:

```go
book := Book{
  Title:    "Kokoro",
  AuthorID: 123,
}
err = booksTable.InsertReturning(&book)

// book should now have an updated ID
```

## Prevent fields from being inserted / updated

If you don't want certaing fields to be inserted or updated if their value is
zero, use the `omitempty` tag on them:

```go
type Book struct {
	ID        uint   `db:"id,omitempty"`
  ...
	AuthorID  uint   `db:"author_id,omitempty"`
	SubjectID uint   `db:"subject_id,omitempty"`
}
```

With `omitempty`, the `id`, `author_id` and `subject_id` fields will be omitted
when they're equal to their zero value.

## The SQL builder

Up to this point you should be able to use CROS, that will work for many of the
tasks on a common CRUD application. Sometimes you'll need to use SQL for tasks
that require it. Continue to our next lesson to learn [how to create SQL
statements using the SQL builder](/tour/06).
