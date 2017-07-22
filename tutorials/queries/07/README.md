# Inserting new items into a collection


Use the `Insert` method on a collection reference to insert a new item:

```go
book := Book{
  Title:    "Kokoro",
  AuthorID: 123,
}

booksTable := sess.Collection("books")

id, err := booksTable.Insert(book)
```

If the table was configured to generate and insert a new primary key
automatically, the `Insert` method will return the ID of the new element (as an
`interface{}`). Note that `Insert` does not modify the passed value at all.

Use the `InsertReturning` method on a collection reference to insert an item
and modify the passed value with fresh values (like IDs or automatic dates)
from the database.

```go
book := Book{
  Title:    "Kokoro",
  AuthorID: 123,
}
err = booksTable.InsertReturning(&book)

// The ID field on the book value should have been updated by now.
log.Printf("New book: %v", book.ID)
```

## Prevent fields from being inserted / updated

If you don't want certain fields to be inserted or updated when their value is
zero use the `omitempty` tag option:

```go
type Book struct {
	ID        uint   `db:"id,omitempty"`

	AuthorID  uint   `db:"author_id,omitempty"`
	SubjectID uint   `db:"subject_id,omitempty"`
}
```

With `omitempty`, the `id`, `author_id` and `subject_id` fields will be omitted
from queries when their values are equal to the zero value of their type.

