package middleware

import web "web-framework-go/v40"

func RecoveryMiddleware() web.HTTPMiddleware {
	return func(next web.HTTPHandleFunc) web.HTTPHandleFunc {
		return func(p7ctx *web.HTTPContext) {
			defer func() {
				if err := recover(); err != nil {
					p7ctx.RespData = append(p7ctx.RespData, []byte("recover from panic\r\n")...)
				}
			}()
			next(p7ctx)
		}
	}
}
