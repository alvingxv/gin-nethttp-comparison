# Gin vs net/http Performance Comparison

This repository contains a performance comparison between the Gin framework and the native net/http package in Go.

To test the performance, use the following `hey` tool command:

```bash
hey -z 10s http://localhost:8000
hey -z 10s http://localhost:8001
```

## Benchmark Results

### Test 1

#### Gin:
- Achieved request/second: 1700~

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	router.Run(":8001")
}
```

#### net/http:
- Achieved request/second: 16000~

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

	http.ListenAndServe(":8000", nil)
}
```

### Test 1

#### Gin:
- Achieved request/second: 16000~

```go
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	s := gin.New()
	s.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	log.Fatal(http.ListenAndServe(":8001", s))
}
```

#### net/http:
- Achieved request/second: 16000~

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

	http.ListenAndServe(":8000", nil)
}
```