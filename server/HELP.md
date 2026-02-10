## Work with Go
- Init project: `go mod init <project-name>`
- Get dependency: `go get <dependency>`
- Clean dependency: `go mod tidy`
- Pull dependency: `go mod download`

## Dependencies
**Web framework**: github.com/gin-gonic/gin
**MongoDB (Official driver)**: go.mongodb.org/mongo-driver/mongo
**Env config**: github.com/joho/godotenv

**Password hashing**: golang.org/x/crypto/bcrypt
**JWT(Json web token)**: github.com/golang-jwt/jwt/v5

**WebSocket**: github.com/gorilla/websocket

**Redis client**: github.com/redis/go-redis/v9

**UUID**: github.com/google/uuid
**Validator(Validate request)**: github.com/go-playground/validator/v10

## Optional
**Defer**
- Defer will run latest in function
- Example:
```go
func PrintText() {
    fmt.Println("1")
    defer fmt.Println("End")
    fmt.Println("2")
}

```
- Output: 1 2 End -> defer line will execute end in function

**Panic**
- Panic will stop execution and throw an error
- Example:
```go
func PrintText() {
    fmt.Println("1")
    panic("Something went wrong")
    fmt.Println("2")
}

```
- Output: 1 panic: Something went wrong -> panic stops execution, line 2 never runs

**Recover**
- Recover will catch panic and prevent program from crashing
- Must be used with defer
- Example:
```go
func PrintText() {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("Caught error:", err)
        }
    }()
    fmt.Println("1")
    panic("Something went wrong")
    fmt.Println("2")
}

```
- Output: 1 Caught error: Something went wrong -> recover catches panic, program continues safely

**Context**
- Context is used to control goroutines and pass values across function boundaries
- Useful for timeout, deadline, and cancellation in concurrent operations
- Example:
```go
func FetchData(ctx context.Context, url string) (string, error) {
    select {
    case <-ctx.Done():
        return "", ctx.Err()  // Context was cancelled
    case result := <-callAPI(url):
        return result, nil
    }
}

// Usage with timeout
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
data, err := FetchData(ctx, "https://api.example.com")

```
- Output: If API takes > 5 seconds, context will timeout and return error -> context prevents hanging requests