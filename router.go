package gwe

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HanderFunc
}

func (r *router) addRoute(method string, pattern string, handler HanderFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}

func newRouter() *router {
	return &router{handlers: map[string]HanderFunc{}}
}
