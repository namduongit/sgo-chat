**Valid, Hash, Etc,...**
- Throw error with type **ServerError**

**MongoDB**
- Throw error with type **MongoError**
- Handle in `exception.go`

**Context in query data**
- If use `context.Background()` in go standard library package
    + Client disconnect -> Query still run
    + Slow query(more data)
    + Can't control time query
- Use `context from HTTP request`
    + Automatically cancel when client disconnects
    + Better control over query timeout
    + Prevents resource waste on abandoned requests