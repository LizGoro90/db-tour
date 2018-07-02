# List All Tables/Collections in a Database

Once `Open` returns a `sess`variable (which satisfies the [db.Database][2] interface), use the `Collections` method to get all the structures in the database: 

```go
collections, err := sess.Collections()
...

log.Printf("Collections in %q: %v", sess.Name(), collections)
```
![Note](https://github.com/LizGoro90/db-tour/tree/master/static/img)
> The db.Database interface provides methods that work on both SQL and NoSQL databases.
> In light of this, the structures of both database types are referred to generally as
> 'collections' and no particular distinction is made between 'tables' and 'collections'.

![Note](https://github.com/LizGoro90/db-tour/tree/master/static/img)
> You can call different db.Database methods depending on the database type. For example, 
> if you're working with a SQL database, `sess` will also satisfy [sqlbuilder.Database][3].


[2]: https://godoc.org/upper.io/db.v3#Database
[3]: https://godoc.org/upper.io/db.v3/lib/sqlbuilder#Database
[4]: https://godoc.org/upper.io/db.v3/lib/sqlbuilder#Open
