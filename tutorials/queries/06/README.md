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

If you want to remove all the items in a result set, use `Delete`:

```go
res := booksTable.Find(4267)

err := res.Delete()
// ...
```

As with the `Update` examples, in the previous case only one item will be affected
and in the following scenario all items will be altered: 


```go
res := booksTable.Find()

err := res.Delete()
...
```

![Note](https://github.com/LizGoro90/db-tour/tree/master/static/img)
> In the particular case of SQL databases, you can also choose to use a builder or 
raw SQL for update, insert, and delete queries. 
