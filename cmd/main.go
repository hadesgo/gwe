package main

import (
	"gwe"
	"log"
	"net/http"
)

func main() {
	r := gwe.New()
	r.GET("/index", func(c *gwe.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(ctx *gwe.Context) {
			ctx.HTML(http.StatusOK, "<h1>Hello Gwe</h1>")
		})

		v1.GET("/hello", func(ctx *gwe.Context) {
			ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("name"), ctx.Path)
		})
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *gwe.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *gwe.Context) {
			c.JSON(http.StatusOK, gwe.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}

	log.Fatal(r.Run(":8000"))
}
