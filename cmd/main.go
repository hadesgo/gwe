package main

import (
	"gwe"
	"net/http"
)

func main() {
	r := gwe.New()
	r.GET("/", func(c *gwe.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *gwe.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *gwe.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *gwe.Context) {
		c.JSON(http.StatusOK, gwe.H{"filepath": c.Param("filepath")})
	})

	r.Run(":8000")
}
