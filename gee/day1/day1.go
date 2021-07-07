package day1

import (
	"fmt"
	"net/http"
)

const address = "localhost:51234"

/*====================================================================================================================*/

/* FirstTryHTTP net.http的用法，监听本机的51234端口，并处理发送到路由/的请求 */
func FirstTryHTTP() {
	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintf(writer, "URL path is %s\n", req.URL.Path)
	})
	_ = http.ListenAndServe(address, nil)
}

/*====================================================================================================================*/

/* Engine1 实现handle接口的ServeHTTP方法，来处理不同路由的请求 */
type Engine1 struct {}

func (e Engine1) ServeHTTP(write http.ResponseWriter, req *http.Request)  {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(write, "URL path is %s\n", req.URL.Path)
	case "/login":
		for k, v := range req.Header {
			fmt.Fprintf(write, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(write, "404 NOT FOUND: %s\n", req.URL)
	}
}

func SecondTryHTTP()  {
	e := &Engine1{}
	_ = http.ListenAndServe(address, e)
}

/*====================================================================================================================*/

/* 接下来实现框架的基本原型，包含路由映射表、启动服务 */
type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine2 struct {
	route map[string]HandlerFunc
}

// New 实例化启动引擎
func New() *Engine2 {
	route := make(map[string]HandlerFunc)
	return &Engine2{
		route: route,
	}
}

// Get 添加Get方法的路由
func (e Engine2) Get(pattern string, handler HandlerFunc) {
	e.addRoute("GET", pattern, handler)
}

// Post 添加Post方法路由
func (e Engine2) Post(pattern string, handler HandlerFunc)  {
	e.addRoute("POST", pattern, handler)
}

func (e Engine2) addRoute(methodType string, pattern string, handler HandlerFunc)  {
	key := combineToRouteKey(methodType, pattern)
	e.route[key] = handler
}

func combineToRouteKey(methodType string, pattern string) string {
	return fmt.Sprintf("%s-%s", methodType, pattern)
}

// ServeHTTP 处理不同路由的请求
func (e Engine2) ServeHTTP(write http.ResponseWriter, req *http.Request)  {
	key := combineToRouteKey(req.Method, req.URL.Path)
	if handler, ok := e.route[key]; ok {
		handler(write, req)
	} else {
		write.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(write, "404 NOT FOUND: %s\n", req.URL)
	}
}

// Run 使用tcp监听给定地址，并启动服务
func (e Engine2) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

/*====================================================================================================================*/

// main day1实现框架的基本使用
func main() {
	s := New()

	s.Get("/gee", func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(writer, "hello, gee")
	})

	s.Run(address)
}



