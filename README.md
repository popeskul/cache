Go in-memory Cache.
=======================

See it in action:

```go
package main

import (
	"fmt"
	"memory-cache/cache"
)

func main() {
	db := map[string]interface{}{}

	cache := cache.New(db)

	cache.Set("userId", 42)
	userId := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId = cache.Get("userId")

	fmt.Println(userId)
}
```
