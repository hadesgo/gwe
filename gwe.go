package gwe

import (
	"net/http"
)

type HanderFunc func(*Context)

type Engine struct {
	router *router
}

func (engine *Engine) addRoute(method string, pattern string, handler HanderFunc) {
	engine.router.addRoute(method, pattern, handler)
}

func (engine *Engine) GET(pattern string, handler HanderFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HanderFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	engine.router.handle(c)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func New() *Engine {
	return &Engine{router: newRouter()}
}
