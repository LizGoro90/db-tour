# Mapping columns to struct fields

Let's suppose we have a table like this:

```sql
CREATE TABLE "books" (
	"id" INTEGER NOT NULL,
	"title" TEXT NOT NULL,
	"author_id" INTEGER,
	"subject_id" INTEGER,
	CONSTRAINT "books_id_pkey" PRIMARY KEY ("id")
);
```

Create a `Book` struct to represent one element from that table and add fields
representing columns:

```go
type Book struct {
	ID          uint
	Title       string
	AuthorID    uint
	SubjectID   uint
}
```

Use the `db` field tag to tell upper-db which database columns
correspond to specific struct fields:

```go
type Book struct {
	ID          uint   `db:"id"`
	Title       string `db:"title"`
	AuthorID    uint   `db:"author_id"`
	SubjectID   uint   `db:"subject_id"`
}
```

Please keep in mind that:

* Fields and columns must be of a similar type, upper-db will handle most
	conversions automatically.
* Fields must be exported, not exported fields will be ignored.

Some databases can be configured to insert automatically-generated values like
IDs, serials, dates or other values if they're not present on a query. If the
table you're working with has a column like that, you'll need to add the
`omitempty` option to its `db` tag:

```go
type Book struct {
	ID uint `db:"id,omitempty"`
}
```

The `omitempty` option will make upper-db ignore zero-valued fields when
building INSERT and UPDATE queries so that they can be correctly generated by
the database itself.

<!--

The following table describes all available options for `db` tags.

| Option           | Description                                     |
| :--------------- | :---------------------------------------------- |
| `omitempty`      | The field is skipped when zero |
-->