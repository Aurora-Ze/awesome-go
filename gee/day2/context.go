package day2

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Context integrate some objects and infos through an interaction process
type Context struct {
	// base objects
	Writer http.ResponseWriter
	Req    *http.Request

	// request info
	Path       string
	Method     string

	// response info
	StatusCode int
}

func newContext(writer http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: writer,
		Req: req,
	}
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// String return format string
func (c *Context) String(code int, format string, values ...interface{})  {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{})  {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}
