package middleware

import (
	"fmt"
	web "web-framework-go/v40"
)

func LogMiddleware() web.HTTPMiddleware {
	return func(next web.HTTPHandleFunc) web.HTTPHandleFunc {
		return func(p7ctx *web.HTTPContext) {
			fmt.Printf("request path:%s\r\n", p7ctx.P7request.URL.Path)
			fmt.Println("ReqBody:", string(p7ctx.ReqBody))
			next(p7ctx)
			fmt.Println("RespData:", string(p7ctx.RespData))
		}
	}
}
