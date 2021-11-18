package gee

import (
	"net/http"
)

//定义函数类型，使实例化了的此类型函数都能加载
type HandleFunc func(c *Context)

//用map存路由映射
type Gee struct {
	router *Router
}

func New() *Gee {
	return &Gee{router: NewRouter()}
}

func (g *Gee) GET(pattern string, handler HandleFunc) {
	g.router.addRouter("GET", pattern, handler)
}

func (g *Gee) POST(pattern string, handler HandleFunc) {
	g.router.addRouter("POST", pattern, handler)
}

func (g *Gee) UPDATE(pattern string, handler HandleFunc) {
	g.router.addRouter("UPDATE", pattern, handler)
}

//gee实现了ServeHTTP方法，可以被实例化为http的Handler
func (g *Gee) Run(addr string) error {
	return http.ListenAndServe(addr, g)
}

//根据http req执行注册的路由
func (g *Gee) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := NewContext(w, req)
	g.router.handle(c)
}
