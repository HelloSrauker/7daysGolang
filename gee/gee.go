package gee

import (
	"log"
	"net/http"
	"strings"
)

//定义函数类型，使实例化了的此类型函数都能加载
type HandleFunc func(c *Context)

//用map存路由映射
type Gee struct {
	*RouterGroup
	router *Router
	groups []*RouterGroup //所有的路由分组
}

type RouterGroup struct {
	prefix      string
	middleWares []HandleFunc // support middleware支持的中间件
	parent      *RouterGroup // support nesting
	gee         *Gee         // all groups share a Engine instance
}

func Default() *Gee {
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}

func New() *Gee {
	gee := &Gee{router: NewRouter()}
	gee.RouterGroup = &RouterGroup{gee: gee}
	gee.groups = []*RouterGroup{gee.RouterGroup}
	return gee
}

//添加路由分组
func (r *RouterGroup) Group(prefix string) *RouterGroup {
	gee := r.gee
	newGroup := &RouterGroup{
		prefix: r.prefix + prefix,
		parent: r,
		gee:    gee,
	}
	gee.groups = append(gee.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) Use(middlewares ...HandleFunc) {
	group.middleWares = append(group.middleWares, middlewares...)
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandleFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.gee.router.addRouter(method, pattern, handler)
}

// GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler HandleFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (group *RouterGroup) POST(pattern string, handler HandleFunc) {
	group.addRoute("POST", pattern, handler)
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
	var middleWares []HandleFunc
	for _, group := range g.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middleWares = append(middleWares, group.middleWares...)
		}
	}

	c := NewContext(w, req)
	c.handlers = middleWares
	g.router.handle(c)
}
