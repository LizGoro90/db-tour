# How to catch common errors

## Enable or disable on-screen logging

Logging is pretty useful when debugging a query. To enable on-screen logging
use:


```go
sess.SetLogging(true)
```

When logging is enabled, upper-db will print queries to stdout. Please keep in
mind that logging is slow and verbose, make sure to disable it on production:

```go
sess.SetLogging(false)
```

## The db.ErrNoMoreRows error

The `db.ErrNoMoreRows` error is returned by `One` or `All` when the result-set
has no items.

```go
// If this table has an integer primary key you can pass an int to Find and
// Find will look for the element that matches that primary key.
err = booksTable.Find(1).One(&book)
if err != nil {
  if err == db.ErrNoMoreRows {
    // This was expected, let's create a new element.
  } else {
    // Something else happened!
    return err
  }
}
```

Depending on your application this error may or may not be fatal.

## Updating and deleting items on a result-set

A result-set is not only for querying data. Result-sets can also be used to
update or delete items. We'll learn more about [how to update and delete
items](/tour/04) on a result-set on our next exercise.
