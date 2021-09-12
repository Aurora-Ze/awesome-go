package test

import (
	"awesome-go/web"
	"fmt"
	"net/http"
	"testing"
	"time"
)

const address = "localhost:51234"

func Test_Web(t *testing.T) {
	engine := web.Default()
	engine.GET("/", func(context *web.Context) {
		time.Sleep(time.Millisecond * 10)
		context.JSON(http.StatusOK, "hello")
	})
	engine.POST("/recovery", func(context *web.Context) {
		arr := []int{1, 2}
		fmt.Println(arr[4])

		context.JSON(http.StatusOK, web.H{
			"msg": "haha",
		})
	})
	engine.POST("/post", func(context *web.Context) {
		body := context.Req.PostForm
		fmt.Printf("body is: \n\t%v\n", body)
	})

	_ = engine.Run(address)
}
