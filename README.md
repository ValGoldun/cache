# Cache

Cache library on pure go

### Example without ttl
```go
cache := cache.New[string, string]()

cache.Set("key","value")

value, found := cache.Get("key")
fmt.Println(value, found) //value true

cache.Delete("key")
```

### Example with ttl
```go
cache := cache.New[string, string]()

cache.SetWithTTL("key","value", time.Minute)

value, found := cache.Get("key")
fmt.Println(value, found) //value true

cache.Delete("key")
```

### Example with default ttl
```go
cache := cache.NewWithTTL[string, string](time.Minute)

cache.Set("key","value")

value, found := cache.Get("key")
fmt.Println(value, found) //value true

cache.Delete("key")
```