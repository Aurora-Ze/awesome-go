package integration

import (
	"helloworld/gee/day5"
	"log"
	"time"
)

// Logger 输出执行一次HTTP请求的相关信息，包括响应码、请求路径和耗时
func Logger() day5.HandlerFunc {
	return func(context *day5.Context) {
		t := time.Now()
		context.Next()
		log.Printf("[%d] %s in %v\n", context.StatusCode, context.Path, time.Since(t))
	}
}
