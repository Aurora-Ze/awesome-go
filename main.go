package main

import (
	"helloworld/gee/day2"
	"net/http"
)

const address = "localhost:51234"

func main() {
	engine := day2.New()

	engine.Get("/gee", func(context *day2.Context) {
		context.JSON(http.StatusOK, []int{1,2,3})
	})

	engine.Post("/gee", func(context *day2.Context) {
		context.JSON(http.StatusOK, day2.H{
			"status": true,
			"geeRsp": &day2.GeeResponse{
				Msg: "hello, welcome to using gee through POST method",
				User: "remiwu",
				Img: nil,
			},
		})
	})
	_ = engine.Run("localhost:51234")
}