package day3

import (
	"fmt"
	"net/http"
	"strings"
)

// router 路由表，支持动态路由，形式为：value 或 *filepath
// 前者表示通配赋值，可用于路径的任意位置；后者表示通配文件路径
// 只能用于路径末尾，且最多出现一个
type router struct {
	roots map[string]*Node // http请求方法与前缀路由树的映射；例如 GET -> a tree，POST -> another tree

	handlers map[string]HandlerFunc // 请求路径与处理方法的映射，key由方法与路径拼接而成；例如 GET-/user/login -> Login method
}

// newRouter 创建路由表
func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
		roots:    make(map[string]*Node),
	}
}

// addRoute 添加路由映射，1.在前缀路由树中添加新路径；2.在路由表中添加一条记录
func (r *router) addRoute(methodType string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)

	root, ok := r.roots[methodType]
	if !ok {
		root = &Node{}
		r.roots[methodType] = root
	}

	root.insert(pattern, parts, 0)
	key := combineToRouteKey(methodType, pattern)
	r.handlers[key] = handler
}

// handle 对请求进行处理
func (r *router) handle(c *Context) {
	// 从路由树中找到匹配的节点
	node, params := r.getRoute(c.Method, c.Path)
	if node != nil {
		c.Params = params // 把解析得到的参数赋值给context，方便后续使用
		key := combineToRouteKey(c.Method, node.pattern)
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}

// getRoute 返回匹配到的节点和参数映射
func (r *router) getRoute(methodType string, pattern string) (*Node, map[string]string) {
	handledParts := parsePattern(pattern)
	root, ok := r.roots[methodType]
	params := make(map[string]string)

	if !ok {
		return nil, nil
	}

	findNode := root.search(handledParts, 0)
	if findNode != nil {
		for i, item := range handledParts {
			if item != "" && item[0] == ':' {
				params[item[1:]] = handledParts[i]
			}
			if item != "" && item[0] == '*' {
				params[item[1:]] = strings.Join(handledParts[i:], "/")
				break
			}
		}
		return findNode, params
	}

	return nil, nil
}

// parsePattern 分割路径, 只取一个*
func parsePattern(pattern string) []string {
	parts := strings.Split(pattern, "/")

	result := make([]string, 0)
	for _, v := range parts {
		if len(v) != 0 {
			result = append(result, v)
			if result[0] == "*" {
				break
			}
		}
	}
	return result
}

// combineToRouteKey 拼接请求方法和路径，形成key
func combineToRouteKey(methodType string, pattern string) string {
	return fmt.Sprintf("%s-%s", methodType, pattern)
}
