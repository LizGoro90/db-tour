# How to update or delete items on a result set

Besides querying and mapping one or many elements into Go values, result-sets
can be also used for updating or deleting all items on the result set.

## Update items on a result-set

If you want to update a result set use the `Update` method:

```go
var book Book
res := booksTable.Find(4267)

err = res.One(&book)
...

err = res.Update(book)
...
```

Remember that `Update` will operate in the whole result set, in this case the
result set has one element.

The example below will update all items in the books table:

```go
res := booksTable.Find()

err := res.Update(map[string]int{
  "author_id": 23,
})
```

## Delete items on a result-set

Like with `Update`, you can delete all items on the result set by using
`Delete`:

```go
res := booksTable.Find(4267)

err := res.Delete()
// ...
```

The example below will delete all items in the books table:


```go
res := booksTable.Find()

err := res.Delete()
...
```

## Insert more items into a collection

Result sets can only work with items that are already part of a collection, if
you want to insert new items you'll have to use the `Insert` or
`InsertReturning` methods on the collection. Continue to our next example to
learn [how to insert new items into a table or collection](/tour/05).
