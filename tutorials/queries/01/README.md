# Map Columns to Struct Fields

Once we get the collections/tables in the database, we'll map them to structs. For 
example, the 'books' table we referred to in the previous step consists of the 
following columns:

```sql
CREATE TABLE "books" (
	"id" INTEGER NOT NULL,
	"title" VARCHAR NOT NULL,
	"author_id" INTEGER,
	"subject_id" INTEGER,
	CONSTRAINT "books_id_pkey" PRIMARY KEY ("id")
);
```

In this case, we need to create a `Book` struct that represents an item from such
table and the fields accompanying it:

```go
type Book struct {
	ID          uint
	Title       string
	AuthorID    uint
	SubjectID   uint
}
```

The `db` field tag will then be required so upper-db can identify the columns
mapped to the struct:

```go
type Book struct {
	ID          uint   `db:"id"`
	Title       string `db:"title"`
	AuthorID    uint   `db:"author_id"`
	SubjectID   uint   `db:"subject_id"`
}
```

![Note](https://github.com/LizGoro90/db-tour/tree/master/static/img)
> Fields and columns must be similar in type (upper-db will handle most conversions
> automatically).

![Note](https://github.com/LizGoro90/db-tour/tree/master/static/img)
> Fields must exported, otherwise they will be ignored.

![Note](https://github.com/LizGoro90/db-tour/tree/master/static/img)
> Use the `db:"-"` tag for exported fields that are not to be mapped.

In the event the table contains a column configured to insert automatically-generated values like IDs, serials, dates, etc. if they're not included in a query, the
`omitempty` option will have to be added to the `db` tag:

```go
type Book struct {
	ID uint `db:"id,omitempty"`
}
```

This option will make upper-db ignore zero-valued fields when
building INSERT and UPDATE queries so they can be correctly generated by
the database itself.

![Note](https://github.com/LizGoro90/db-tour/tree/master/static/img)
> You can use different mapping methods depending on the database type. To learn more,
click [here](https://upper.io/db.v3/getting-started#mapping-tables-to-structs).

<!--

The following table describes all available options for `db` tags.

| Option           | Description                                     |
| :--------------- | :---------------------------------------------- |
| `omitempty`      | The field is skipped when zero |
-->
