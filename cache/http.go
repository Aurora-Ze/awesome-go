package cache

import (
	"awesome-go/web"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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

func StartClient(group, key string) ([]byte, error) {
	_ = NewHttpPool("data1")
	rsp, err := http.Get(Address + DefaultBaseURL + url.QueryEscape(group) + url.QueryEscape(key))
	if err != nil {
		log.Printf("[Distributed Cache] http error\n")
		return nil, err
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		log.Printf("[Distributed Cache] status code %d", rsp.StatusCode)
		return nil, fmt.Errorf("status code error, code = %d\n", rsp.StatusCode)
	}

	bytes, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Printf("[Distributed Cache] read error, error = %v\n", err)
		return nil, fmt.Errorf("io read error, error = %v\n", err)
	}

	return bytes, nil
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
