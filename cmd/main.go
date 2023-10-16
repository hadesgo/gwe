package main

import (
	"gwe"
	"log"
	"net/http"
	"time"
)

func onlyForV2() gwe.HandlerFunc {
	return func(ctx *gwe.Context) {
		t := time.Now()
		ctx.Fail(500, "Internal Server Error")
		log.Printf("[%d] %s in %v for group v2", ctx.StatusCode, ctx.Request.RequestURI, time.Since(t))
	}
}

func main() {
	r := gwe.New()
	r.Use(gwe.Logger())

	r.GET("/", func(c *gwe.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *gwe.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	log.Fatal(r.Run(":8000"))
}
