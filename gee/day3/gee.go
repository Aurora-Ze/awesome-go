package day3

import "net/http"

type HandlerFunc func(*Context)

type Engine struct {
	*router
}

type H map[string]interface{}

// New this method create and return a route engine
func New() *Engine {
	return &Engine{
		newRouter(),
	}
}

func (e Engine) addRoute(methodType string, pattern string, handler HandlerFunc) {
	key := combineToRouteKey(methodType, pattern)
	e.router.handlers[key] = handler
}

// Get add a new record to router tables, which contains a GET method with its path and handler
func (e Engine) Get(pattern string, handler HandlerFunc) {
	e.addRoute("GET", pattern, handler)
}

// Post add a new record to router tables, which contains a POST method with its path and handler
func (e Engine) Post(pattern string, handler HandlerFunc) {
	e.addRoute("POST", pattern, handler)
}

// ServeHTTP handle the request with context
func (e Engine) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	context := newContext(writer, req)
	context.Method = req.Method
	context.Path = req.URL.Path
	e.router.handle(context)
}

// Run start server
func (e Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

/*====================================================================================================================*/

type GeeResponse struct {
	Msg  string
	User string
	Img  []string
}

func main() {
	engine := New()

	engine.Get("/gee", func(context *Context) {
		context.JSON(http.StatusOK, &GeeResponse{
			Msg:  "hello, welcome to using gee",
			User: "aurora",
			Img:  nil,
		})
	})

	engine.Post("/gee", func(context *Context) {
		context.JSON(http.StatusOK, H{
			"status": true,
			"geeRsp": &GeeResponse{
				Msg:  "hello, welcome to using gee through POST method",
				User: "remiwu",
				Img:  nil,
			},
		})
	})
	_ = engine.Run("localhost:51234")
}
