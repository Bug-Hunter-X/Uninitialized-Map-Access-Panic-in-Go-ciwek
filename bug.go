```go
func main() {
    var m map[string]int
    m["key"] = 10 // This will panic if the map is not initialized
    fmt.Println(m["key"])
}
```