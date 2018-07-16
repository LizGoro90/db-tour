# Update, Insert, or Delete Items in a Result Set

Besides querying for data, result-sets can also be used for updating or
deleting all matching items.

If you want to update all the items on a result-set use the `Update` method:

```go
var book Book
res := booksTable.Find(4267)

err = res.One(&book)
...

err = res.Update(book)
...
```

Remember that `Update` will operate in the whole result-set. In the example
above the whole result-set consists of only one element, while in the example
below the result-set consists of all the items in the collection:

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
