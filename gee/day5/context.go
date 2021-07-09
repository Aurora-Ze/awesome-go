package day5

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Context 集成了一次HTTP请求与响应过程中的一些数据，包括 Method、 Path、 Params 和 StatusCode 等
type Context struct {
	// base objects
	Writer http.ResponseWriter
	Req    *http.Request

	// request info
	Path   string
	Method string
	Params map[string]string
	// response info
	StatusCode int
	// middlewares
	handlers []HandlerFunc
	index    int
}

func newContext(writer http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: writer,
		Req:    req,
		Method: req.Method,
		Path:   req.URL.Path,
		index:  -1,
	}
}

func (c *Context) Next() {
	c.index++

	len := len(c.handlers)
	for ; c.index < len; c.index++ {
		c.handlers[c.index](c)
	}
}

// Status 设置HTTP响应状态码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader 设置HTTP响应头键值对
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// Param 返回 Context 中得到的参数值
func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

// String 返回格式化字符串
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON 返回 JSON 格式
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// PostForm 获取Post请求携带的实体的值
func (c *Context) PostForm(key string) string {
	_ = c.Req.ParseForm()
	return c.Req.FormValue(key)
}

// Query 表示路径后面的参数，例如/login?name=zhangsan
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}
