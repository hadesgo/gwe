package main

import (
	"gwe"
	"net/http"
)

func main() {
	engine := gwe.New()
	engine.GET("/", func(ctx *gwe.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello Gwe</h1>")
	})
	engine.GET("/hello", func(ctx *gwe.Context) {
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("name"), ctx.Path)
	})
	engine.POST("/login", func(ctx *gwe.Context) {
		ctx.JSON(http.StatusOK, gwe.H{
			"username": ctx.PostForm("username"),
			"password": ctx.PostForm("password"),
		})
	})
	engine.Run(":8000")
}
