package day3

import "net/http"

type HandlerFunc func(*Context)

type Engine struct {
	*router
}

type H map[string]interface{}

// New 返回一个Engine实例
func New() *Engine {
	return &Engine{
		newRouter(),
	}
}

// ServeHTTP 把信息封装成Context，并调用router去处理请求。每次请求都会调用该方法。
func (e Engine) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	context := newContext(writer, req)
	context.Method = req.Method
	context.Path = req.URL.Path
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
	return http.ListenAndServe(addr, e)
}

// addRoute 添加路由映射
func (e Engine) addRoute(methodType string, pattern string, handler HandlerFunc) {
	key := combineToRouteKey(methodType, pattern)
	e.router.handlers[key] = handler
	e.router.addRoute(methodType, pattern, handler)
}
