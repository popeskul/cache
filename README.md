Go in-memory Cache.
=======================

See it in action:

```go
package main

import (
	"fmt"
	"github.com/popeskul/cache"
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

![Go Report](https://goreportcard.com/badge/github.com/popeskul/cache)
![Repository Top Language](https://img.shields.io/github/languages/top/popeskul/cache)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/popeskul/cache)
![Github Repository Size](https://img.shields.io/github/repo-size/popeskul/cache)
![Github Open Issues](https://img.shields.io/github/issues/popeskul/cache)
![Lines of code](https://img.shields.io/tokei/lines/github/popeskul/cache)
![License](https://img.shields.io/badge/license-MIT-green)
![GitHub last commit](https://img.shields.io/github/last-commit/popeskul/cache)