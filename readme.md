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

### Test 2

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

## Summary:

In Gin, the default mode uses the net/http package under the hood. However, for the best performance, it is recommended to use gin.New() instead of gin.Default().

When using gin.Default(), Gin sets up some default middleware handlers, such as logging and recovery, which provide convenience but add some overhead to the request handling process. These default middleware handlers are useful for development and debugging purposes.

On the other hand, gin.New() creates a new Gin engine without any default middleware handlers. This allows for a leaner and faster setup since you have full control over which middleware handlers to add, resulting in improved performance.

By using gin.New(), you can choose and add only the necessary middleware handlers based on your specific requirements, optimizing the performance of your Gin application.

Overall, while gin.Default() provides convenience by including default middleware handlers, if performance is a top priority, it is recommended to use gin.New() and selectively add middleware handlers as needed.
