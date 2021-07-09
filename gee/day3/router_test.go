package day3

import (
	"log"
	"testing"
)

func testNewRouter() *router {
	route := newRouter()
	route.addRoute("GET", "/user/login/", func(context *Context) {})
	route.addRoute("GET", "user/login1/:userID", func(context *Context) {})
	route.addRoute("GET", "user/login/*filepath", func(context *Context) {})

	return route

	//node, _ := route.roots["GET"]
	//fmt.Printf("%+v\n", node.children[0].children[0].children[0])
}

func Test_getRoute(t *testing.T) {
	route := testNewRouter()

	node, params := route.getRoute("GET", "/user/login")

	if node == nil {
		log.Fatal("getRoute find empty")
	}
	if len(params) == 0 {
		log.Println("params empty")
	}
	log.Printf("node is %+#v\n", node)

}
