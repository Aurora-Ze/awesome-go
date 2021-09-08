package web

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
	root := newRouterGroup()
	engine := &Engine{
		root,
		newRouter(),
		[]*RouterGroup{root},
	}
	root.engine = engine

	return engine
}

// Default 返回一个默认的Engine，使用Logger和Recovery中间件
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

// Run 开启服务
func (e Engine) Run(addr string) error {
	err := http.ListenAndServe(addr, e)

	if err != nil {
		log.Fatalf("Start serve error: %v", err)
	}

	log.Printf("Start serve, listen on %s", addr)
	return nil
}
