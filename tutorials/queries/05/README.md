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

Foreseeing solutions that resume execution when it is interrupted due to an
error is a good practice too. In the case below, where int is passed to Find
to look for an integer primary key in the "books" table, different error 
scenarios can be defined. For example, `db.ErrNoMoreRows`, which is returned 
by `One` or `All` when the result set has zero items.

```go
err = booksTable.Find(1).One(&book)
if err != nil {
  if err == db.ErrNoMoreRows {
    // First possible error scenario
  } else {
    // All other possible error scenarios
    return err
  }
}
```

![Note](https://github.com/LizGoro90/db-tour/tree/master/static/img)
> Error scenarios may or may not be fatal depending on the nature of your
application, so make sure you're handling them properly. 

![Note](https://github.com/LizGoro90/db-tour/tree/master/static/img)
> The db.v3 package defines different variables that return error messages.
To learn more, click [here](https://godoc.org/upper.io/db.v3#pkg-variables)
