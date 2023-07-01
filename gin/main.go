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
