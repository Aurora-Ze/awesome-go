package day5

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
	"time"
)

// 声明各种类型的中间件。如果想要在请求执行前添加一些功能，
// 请把相应代码放在调用context.Next()之前；反之，则把代码放在调用context.Next()之后。

// Logger 输出执行一次HTTP请求的相关信息，包括响应码、请求路径和耗时
func Logger() HandlerFunc {
	return func(context *Context) {
		t := time.Now()
		context.Next()
		log.Printf("[%d] %s in %v ms\n", context.StatusCode, context.Path, time.Since(t).Milliseconds())
	}
}

// Recovery 错误恢复，发生panic时尝试进行恢复，
func Recovery() HandlerFunc {
	return func(context *Context) {
		defer func() {
			// 由recover返回值来确定当前是否有panic发生
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				context.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()

		context.Next()
	}
}

// trace 打印错误的堆栈信息
func trace(message string) string {
	var pcs [32]uintptr
	count := runtime.Callers(3, pcs[:])

	var builder strings.Builder
	builder.WriteString(message + "\nTraceback:\n")
	for _, pc := range pcs[:count] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		builder.WriteString(fmt.Sprintf("\t%s:%d\n", file, line))
	}

	return builder.String()
}
