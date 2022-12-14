package main

import (
	"fmt"
	"net/http"
	webframework "web-framework-go/v22"
)

type ApiJson struct {
	JsonInt    int    `json:"json_int"`
	JsonString string `json:"json_string"`
	JsonText   string `json:"json_text"`
}

func main() {
	p1hshutdown := webframework.NewHTTPShutdown()

	p1hservice := webframework.NewHTTPSrevice(
		"http-service",
		p1hshutdown.ReqInHandleCountBuilder,
		webframework.TestMiddlewareBuilder,
		webframework.TimeCostMiddlewareBuilder,
	)

	httpApi(p1hservice)
	go p1hservice.Start("127.0.0.1", "9501")

	webframework.WaitForShutdown(
		webframework.NotifyShutdownToGateway,
		p1hshutdown.RejectNewRequestAndWaiting,
		// 全部请求处理完了，就可以关闭服务了
		webframework.ServiceShutdownBuilder(p1hservice))

	fmt.Println("done")
}

func httpApi(p1hservice webframework.Service) {
	p1hservice.RegisteRoute(http.MethodGet, "/api/test", func(p1c *webframework.HTTPContext) {
		p1c.P1resW.WriteHeader(http.StatusOK)
		_, _ = p1c.P1resW.Write([]byte("response, http.MethodGet, /api/test"))
	})

	p1hservice.RegisteRoute(http.MethodPost, "/api/post_json", func(p1c *webframework.HTTPContext) {
		reqData := &ApiJson{}
		err := p1c.ReadJson(reqData)
		if nil != err {
			p1c.WriteJson(http.StatusUnprocessableEntity, err.Error())
			return
		}
		reqData.JsonText = "response, http.MethodPost, /api/json"
		p1c.WriteJson(http.StatusOK, reqData)
	})

	p1hservice.RegisteRoute(http.MethodGet, "/user/info", func(p1c *webframework.HTTPContext) {
		p1c.P1resW.WriteHeader(http.StatusOK)
		_, _ = p1c.P1resW.Write([]byte("response, http.MethodGet, /user/info/1"))
	})

	p1hservice.RegisteRoute(http.MethodGet, "/user/*", func(p1c *webframework.HTTPContext) {
		p1c.P1resW.WriteHeader(http.StatusOK)
		_, _ = p1c.P1resW.Write([]byte("response, http.MethodGet, /user/*"))
	})

	p1hservice.RegisteRoute(http.MethodGet, "/user/:id", func(p1c *webframework.HTTPContext) {
		p1c.P1resW.WriteHeader(http.StatusOK)
		_, _ = p1c.P1resW.Write([]byte("response, http.MethodGet, /user/" + p1c.PathParams["id"]))
	})

	p1hservice.RegisteRoute(http.MethodGet, "/user/order", func(p1c *webframework.HTTPContext) {
		p1c.P1resW.WriteHeader(http.StatusOK)
		_, _ = p1c.P1resW.Write([]byte("response, http.MethodGet, /user/order"))
	})

	p1hservice.RegisteRoute(http.MethodGet, "/user/order/:id/detail", func(p1c *webframework.HTTPContext) {
		p1c.P1resW.WriteHeader(http.StatusOK)
		_, _ = p1c.P1resW.Write([]byte("response, http.MethodGet, /user/order/" + p1c.PathParams["id"] + "/detail"))
	})
}
