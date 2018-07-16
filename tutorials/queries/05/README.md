# Debug

Logging is pretty useful for finding errors in your code. To enable it, use:

```go
sess.SetLogging(true)
```

When logging is enabled, upper-db will print queries to `stdout`. Please keep
in mind that logging is slow and verbose, make sure to disable it on
production:

```go
sess.SetLogging(false)
```

## Error handling

The `db.ErrNoMoreRows` error is returned by `One` or `All` when the result-set
has zero items.

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

Depending on your application this error may or may not be fatal, make sure
you're handling it properly.
