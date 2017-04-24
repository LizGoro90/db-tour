# How to use transactions

Use the `Tx` method on a SQL database session to create a transaction block,
this method expects a context and a function (`func(sqlbuilder.Tx) error`).

```go
ctx := context.Background()

err := sess.Tx(ctx, func(tx sqlbuilder.Tx) error {
  ...
})
```

The `ctx` value can be used to cancel and rollback a transaction before it
ends. The transaction function defines what you want to do within a transaction
context and receives a ready-to-be-used transaction session `tx`. This `tx`
value can be used like a regular `sess` except that it actually is a
transaction.

If the passed function returns an error the transaction is rolled back:

```go
err := sess.Tx(ctx, func(tx sqlbuilder.Tx) error {
  ...
  return errors.New("Transaction failed")
})
```

If the passed function returns `nil`, the transaction will be commited.

```go
err := sess.Tx(ctx, func(tx sqlbuilder.Tx) error {
  ...
  return nil
})
```

## More questions?

This is the end of the upper-db tutorial, if you have any other questions
feel free to ask on the #upper-db channel at...
