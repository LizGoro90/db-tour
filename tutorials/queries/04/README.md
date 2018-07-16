# Paginate Results

The pagination API lets you split the results of a query into chunks containing a
fixed number of items.

### Simple Pagination 

You can use numbered pages, for example:

```go
// Create paginator and set the amount of items by chunk
res = sess.Collection("posts").Paginate(20) 

// Get first chunk of results (page 1)
err = res.All(&posts) 

// Get second chunk of results (page 2)(limit 20, offset 40)
err = res.Page(2).All(&posts) 
```

![Note](https://github.com/LizGoro90/db-tour/tree/master/static/img)
> If you're working with SQL builder, use `SelectFrom` instead of Collection: 

```go
q = sess.SelectFrom("posts").Paginate(20)
```

### Cursor-based Pagination 

You can also set the item where you want to begin and the results you want to 
fetch thereon:

```go
res = sess.Collection("posts").
  Paginate(20). // 20 results per page
  Cursor("id") // using id as cursor

err = res.All(&posts) // First 20 results of the query

// Get the next 20 results starting from the last item of the previous query.
res = res.NextPage(posts[len(posts)-1].ID)
err = res.All(&posts) // Results from page 1, limit 20, offset 20
```

### Other commonly used pagination tools

```go
res = res.Paginate(23)

totalNumberOfEntries, err = res.TotalEntries()

totalNumberOfPages, err = res.TotalPages()
```

