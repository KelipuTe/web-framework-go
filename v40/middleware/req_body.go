package middleware

import (
	"bytes"
	"io"
	web "web-framework-go/v40"
)

// ReqBodyMiddleware 提取请求参数
func ReqBodyMiddleware() web.HTTPMiddleware {
	return func(next web.HTTPHandleFunc) web.HTTPHandleFunc {
		return func(p7ctx *web.HTTPContext) {
			// 处理请求参数
			var err error
			p7ctx.ReqBody, err = io.ReadAll(p7ctx.P7request.Body)
			if nil != err {
				return
			}
			p7ctx.P7request.Body = io.NopCloser(bytes.NewBuffer(p7ctx.ReqBody))

			next(p7ctx)
		}
	}
}
