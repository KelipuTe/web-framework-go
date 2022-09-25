package main

import (
	"fmt"
	"net/http"
	"time"
	web "web-framework-go/v40"
	"web-framework-go/v40/middleware"
	"web-framework-go/v40/shutdown"
)

func main() {
	p7os := makeOpenService()
	p7as := makeAdminService()
	p7sm := web.NewServiceManager(
		[]*web.HTTPService{p7os, p7as},
		web.SetShutdownTimeOutOption(20*time.Second),
		web.SetShutdownWaitTime(10*time.Second),
		web.SetShutdownCallbackTimeOut(5*time.Second),
	)
	p7sm.Start()
}

func makeOpenService() *web.HTTPService {
	p7h := web.NewHTTPHandler()

	p7h.AddMiddleware(
		middleware.RecoveryMiddleware(),
		middleware.ReqBodyMiddleware(),
		middleware.LogMiddleware(),
	)

	f4handler := func(p7ctx *web.HTTPContext) {
		routingInfo := p7ctx.GetRoutingInfo()
		pathParam := "pathParam:"
		for key, val := range p7ctx.M3pathParam {
			pathParam += fmt.Sprintf("%s=%s;", key, val)
		}
		p7ctx.RespData = append(p7ctx.RespData, []byte(routingInfo+pathParam)...)
	}

	p7h.Get("/", f4handler)

	p7h.Get("/hello", f4handler)
	p7h.Get("/hello/world", f4handler, middleware.TestMiddleware("/hello"), middleware.TestMiddleware("/world"))
	p7h.Get("/hello/*", f4handler, middleware.TestMiddleware("/hello/*"))

	p7h.Get("/order", f4handler)
	p7h.Get("/order/list/:size/:page", f4handler)
	p7h.Get("/order/:id/detail", f4handler)
	p7h.Post("/order/create", f4handler)
	p7h.Post("/order/:id/delete", f4handler)

	p7s := web.NewHTTPService("9510", "127.0.0.1:9510", p7h)

	p7s.AddShutdownCallback(
		shutdown.CacheShutdownCallback,
		shutdown.CountShutdownCallback,
	)

	return p7s
}

func makeAdminService() *web.HTTPService {
	p7h := web.NewHTTPHandler()

	p7h.AddMiddleware(
		middleware.RecoveryMiddleware(),
		middleware.ReqBodyMiddleware(),
		middleware.LogMiddleware(),
	)

	f4handler := func(p7ctx *web.HTTPContext) {
		routingInfo := p7ctx.GetRoutingInfo()
		pathParam := "pathParam:"
		for key, val := range p7ctx.M3pathParam {
			pathParam += fmt.Sprintf("%s=%s;", key, val)
		}
		p7ctx.RespData = append(p7ctx.RespData, []byte(routingInfo+pathParam)...)
	}

	p7h.Group(
		"/admin",
		[]web.HTTPMiddleware{middleware.TestMiddleware("admin")},
		[]web.RouteData{
			{http.MethodGet, "/", f4handler},
			{http.MethodGet, "/list/:size/:page", f4handler},
			{http.MethodGet, "/:id/detail", f4handler},
			{http.MethodPost, "/create", f4handler},
			{http.MethodPost, "/:id/delete", f4handler},
		},
	)

	return web.NewHTTPService("9511", "127.0.0.1:9511", p7h)
}
