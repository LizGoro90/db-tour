# Debug

Logging is pretty useful for finding errors in your code. To enable it, use:

```go
sess.SetLogging(true)
```

This way upper-db will print queries to `stdout`. 

![Note](https://github.com/LizGoro90/db-tour/tree/master/static/img)
> Make sure to disable logging in production as it is slow and verbose. 

```go
sess.SetLogging(false)
```

# Handle Errors

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
