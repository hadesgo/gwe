package main

import (
	"gwe"
	"log"
	"net/http"
)

func main() {
	r := gwe.Default()
	r.GET("/", func(c *gwe.Context) {
		c.String(http.StatusOK, "Hello Geektutu\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *gwe.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})

	log.Fatal(r.Run(":8000"))
}
