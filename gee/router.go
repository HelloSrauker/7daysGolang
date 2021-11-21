package gee

import (
	"net/http"
)

type Router struct {
	handler map[string]HandleFunc
}

func NewRouter() *Router {
	return &Router{handler: map[string]HandleFunc{}}
}

//将路由存进路由映射
func (r *Router) addRouter(method, pattern string, handler HandleFunc) {
	key := method + "_" + pattern
	r.handler[key] = handler
}

func (r *Router) handle(c *Context) {
	key := c.Method + "_" + c.Path
	if handler, ok := r.handler[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND : %+v", key)
	}
}
