package day2

import (
	"fmt"
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	handlers := make(map[string]HandlerFunc)
	return &router{
		handlers: handlers,
	}
}

func (r *router) addRoute(methodType string, pattern string, handler HandlerFunc)  {
	key := combineToRouteKey(methodType, pattern)
	r.handlers[key] = handler
}

func (r *router) handle(c *Context)  {
	key := combineToRouteKey(c.Req.Method, c.Req.URL.Path)
	fmt.Print(key)
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}

func combineToRouteKey(methodType string, pattern string) string {
	return fmt.Sprintf("%s-%s", methodType, pattern)
}
