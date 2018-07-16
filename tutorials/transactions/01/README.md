### Transactions

Use the `Tx` method on a SQL database session to create a transaction block,
this method expects a `context.Context` value and a `func(sqlbuilder.Tx) error`
function.

```go
ctx := context.Background()

err := sess.Tx(ctx, func(tx sqlbuilder.Tx) error {
  ...
})
```

The `ctx` value can be used to cancel and rollback a transaction before it
ends. The transaction function defines what you want to do within a transaction
context and receives a ready-to-be-used transaction session `tx`. This `tx`
value can be used like a regular `sess` except that any write operation that
happens on it needs to be either commited or rolled back.

If the passed function returns an error the transaction gets rolled back:

```go
err := sess.Tx(ctx, func(tx sqlbuilder.Tx) error {
  ...
  return errors.New("Transaction failed")
})
```

If the passed function returns `nil`, the transaction gets commited.

```go
err := sess.Tx(ctx, func(tx sqlbuilder.Tx) error {
  ...
  return nil
})
```

<!--
## More questions?

This is the end of the upper-db tutorial! if you have any more questions
feel free to ask on the #upper-db channel at...
-->
