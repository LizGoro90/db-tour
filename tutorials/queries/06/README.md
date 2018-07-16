# Update, Insert, or Delete Items in a Result Set

The items in result sets can not only be queried, but also modified and removed.

If you want to modify the properties of a complete result set, use `Update`:

```go
var book Book
res := booksTable.Find(4267)

err = res.One(&book)
...

err = res.Update(book)
...
```

Note that this result set consists of only one element, whereas the next result
set consists of all the items in the collection:

```go
res := booksTable.Find()

err := res.Update(map[string]int{
  "author_id": 23,
})
```

As with `Update`, you can delete all items on the result set by using `Delete`:

```go
res := booksTable.Find(4267)

err := res.Delete()
// ...
```

The example above affects one item, and the example below will delete all items
in the books table:


```go
res := booksTable.Find()

err := res.Delete()
...
```
