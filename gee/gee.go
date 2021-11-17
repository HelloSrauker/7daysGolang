package gee

import (
	"fmt"
	"net/http"
)

//定义函数类型，使实例化了的此类型函数都能加载
type HandleFunc func(http.ResponseWriter, *http.Request)

//用map存路由映射
type Gee struct {
	router map[string]HandleFunc
}

func New() *Gee {
	return &Gee{router: map[string]HandleFunc{}}
}

//将路由存进路由映射
func (g *Gee) addRouter(method, pattern string, handler HandleFunc) {
	key := method + "_" + pattern
	g.router[key] = handler
}

func (g *Gee) GET(pattern string, handler HandleFunc) {
	g.addRouter("GET", pattern, handler)
}

func (g *Gee) POST(pattern string, handler HandleFunc) {
	g.addRouter("POST", pattern, handler)
}

func (g *Gee) UPDATE(pattern string, handler HandleFunc) {
	g.addRouter("UPDATE", pattern, handler)
}

//gee实现了ServeHTTP方法，可以被实例化为http的Handler
func (g *Gee) Run(addr string) error {
	return http.ListenAndServe(addr, g)
}

//根据http req执行注册的路由
func (g *Gee) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "_" + req.URL.Path
	if handler, ok := g.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND : %+v", key)
	}
}
