package test

import (
	"fmt"
	"helloworld/gee/day2"
	"helloworld/gee/day3"
	"helloworld/gee/day4"
	"helloworld/gee/day5"
	"net/http"
	"time"
)

const address = "localhost:51234"

func TestDay2() {
	engine := day2.New()

	engine.Get("/gee", func(context *day2.Context) {
		context.JSON(http.StatusOK, []int{1, 2, 3})
	})

	engine.Post("/gee", func(context *day2.Context) {
		context.JSON(http.StatusOK, day2.H{
			"status": true,
			"geeRsp": &day2.GeeResponse{
				Msg:  "hello, welcome to using gee through POST method",
				User: "remiwu",
				Img:  nil,
			},
		})
	})
	_ = engine.Run(address)
}

func TestDay3() {
	engine := day3.New()

	engine.Get("/gee", func(context *day3.Context) {
		context.JSON(http.StatusOK, day3.H{
			"gee": "hello",
		})
	})
	engine.Post("/gee", func(context *day3.Context) {
		context.JSON(http.StatusOK, day3.H{
			"gee": context.PostForm("gee"),
		})
	})

	engine.Run(address)
}

func TestDay4() {
	engine := day4.New()

	rg := engine.Group("/user")
	rg.GET("/login/:name", func(context *day4.Context) {
		context.JSON(http.StatusOK, day4.H{
			"name": context.Param("name"),
		})
	})

	_ = engine.Run(address)
}

func TestDay5() {
	engine := day5.Default()
	engine.GET("/", func(context *day5.Context) {
		time.Sleep(time.Millisecond * 10)
		context.JSON(http.StatusOK, "hello")
	})
	engine.POST("/recovery", func(context *day5.Context) {
		arr := []int{1, 2}
		fmt.Println(arr[4])

		context.JSON(http.StatusOK, day5.H{
			"msg": "haha",
		})
	})
	engine.POST("/post", func(context *day5.Context) {
		body := context.Req.PostForm
		fmt.Printf("body is: \n\t%v\n", body)
	})

	engine.Run(address)
}
