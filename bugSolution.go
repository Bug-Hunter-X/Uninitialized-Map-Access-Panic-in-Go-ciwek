```go
package main

import "fmt"

func main() {
    var m map[string]int
    m = make(map[string]int)
    m["key"] = 10 
    fmt.Println(m["key"])
}
```