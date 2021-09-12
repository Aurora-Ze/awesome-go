package cache

import (
	"awesome-go/web"
	"log"
	"net/http"
)

const (
	DefaultBaseURL = "/cache_api/"
	Address        = "localhost:9999"
)

type HttpPool struct {
	self    string
	baseURL string
}

// todo not used
func NewHttpPool(self string) *HttpPool {
	p := &HttpPool{
		self:    self,
		baseURL: DefaultBaseURL,
	}
	return p
}

func StartServer() {
	engine := web.New()
	engine.GET("/cache_api/:group/:key", ServerHandler)
	_ = engine.Run(Address)
}

func ServerHandler(context *web.Context) {
	// resolve params
	groupName, ok1 := context.Params["group"]
	key, ok2 := context.Params["key"]
	if !ok1 || !ok2 {
		log.Printf("can not found groupName or key")
		return
	}

	group := GetGroup(groupName)
	if group == nil {
		log.Printf("group \"%s\" not exist\n", groupName)
		return
	}

	data, err := group.Get(key)
	if err != nil {
		log.Printf("error occurs, err = %v\n", err)
		return
	}

	context.String(http.StatusOK, "data = %s\n", data.String())
}
