package day5

import (
	"log"
	"net/http"
	"strings"
)

type HandlerFunc func(*Context)

// Engine 启动引擎，嵌套了路由组，因此具有路由组的所有方法。
type Engine struct {
	*RouterGroup                // 根路由组
	router       *router        // 路由表
	groups       []*RouterGroup // 持有所有路由组
}

type H map[string]interface{}

// New 返回一个Engine实例
func New() *Engine {
	rootRouter := newRouterGroup()
	engine := &Engine{
		rootRouter,
		newRouter(),
		[]*RouterGroup{rootRouter},
	}
	rootRouter.engine = engine

	return engine
}

func Default() *Engine {
	e := New()
	e.Use(Logger(), Recovery())
	return e
}

// ServeHTTP 把信息封装成Context，并调用router去处理请求。每次请求都会调用该方法。
func (e Engine) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	handler := make([]HandlerFunc, 0)
	for _, group := range e.groups {
		if strings.HasPrefix(req.URL.Path, group.baseURL) {
			handler = append(handler, group.handler...)
		}
	}

	context := newContext(writer, req)
	context.handlers = handler
	e.router.handle(context)
}

// Get 添加一个 GET 方法的路由
func (e Engine) Get(pattern string, handler HandlerFunc) {
	e.addRoute("GET", pattern, handler)
}

// Post 添加一个 POST 方法的路由
func (e Engine) Post(pattern string, handler HandlerFunc) {
	e.addRoute("POST", pattern, handler)
}

// Run 开启服务
func (e Engine) Run(addr string) error {
	err := http.ListenAndServe(addr, e)

	if err != nil {
		log.Fatalf("Start serve error: %v", err)
	}

	log.Printf("Start serve, listen on %s", addr)
	return nil
}

// addRoute 添加路由映射
func (e Engine) addRoute(methodType string, pattern string, handler HandlerFunc) {
	key := combineToRouteKey(methodType, pattern)
	e.router.handlers[key] = handler
	e.router.addRoute(methodType, pattern, handler)
}
