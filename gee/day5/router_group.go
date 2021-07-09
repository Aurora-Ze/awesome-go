package day5

// RouterGroup 路由组，方便对一组路由进行管理和配置。所有路由组拥有一个共同的 Engine
type RouterGroup struct {
	handler []HandlerFunc //
	parent  *RouterGroup  // 父路由组
	baseURL string        // 本组路由的 path
	engine  *Engine
}

func newRouterGroup() *RouterGroup {
	return &RouterGroup{}
}

// Use 为当前路由组添加中间件
func (r *RouterGroup) Use(handler ...HandlerFunc) {
	r.handler = append(r.handler, handler...)
}

// Group 声明一个路由组
func (r *RouterGroup) Group(baseURL string) *RouterGroup {
	group := &RouterGroup{
		baseURL: r.baseURL + baseURL,
		engine:  r.engine,
		parent:  r,
	}

	r.engine.groups = append(r.engine.groups, group)
	return group
}

func (r *RouterGroup) GET(pattern string, handler HandlerFunc) {
	r.addRoute("GET", pattern, handler)
}

func (r *RouterGroup) POST(pattern string, handler HandlerFunc) {
	r.addRoute("POST", pattern, handler)
}

func (r *RouterGroup) addRoute(methodType string, pattern string, handler HandlerFunc) {
	r.engine.router.addRoute(methodType, r.baseURL+pattern, handler)
}
