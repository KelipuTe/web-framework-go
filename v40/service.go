package v40

import (
	"fmt"
	"net/http"
)

// HTTPHandleFunc 路由对应的处理方法的定义
type HTTPHandleFunc func(p7ctx *HTTPContext)

// HTTPServiceInterface 核心服务的接口定义
type HTTPServiceInterface interface {
	http.Handler
	Start(addr string) error
	RouterInterface
}

// HTTPService 核心服务
type HTTPService struct {
	router
	s5f4middleware []HTTPMiddleware
}

// 确保 HTTPService 实现了 HTTPServiceInterface 接口
var _ HTTPServiceInterface = &HTTPService{}

func NewHTTPService() *HTTPService {
	return &HTTPService{
		router: newRouter(),
	}
}

// Get 包装 addRoute
func (p7this *HTTPService) Get(path string, f4h HTTPHandleFunc) {
	p7this.router.addRoute(http.MethodGet, path, f4h)
}

// Post 包装 addRoute
func (p7this *HTTPService) Post(path string, f4h HTTPHandleFunc) {
	p7this.router.addRoute(http.MethodPost, path, f4h)
}

func (p7this *HTTPService) ServeHTTP(i9w http.ResponseWriter, p7r *http.Request) {
	p7ctx := &HTTPContext{
		I9writer:  i9w,
		P7request: p7r,
	}
	p7this.doServeHTTP(p7ctx)
}

func (p7this *HTTPService) doServeHTTP(p7ctx *HTTPContext) {
	p7ri := p7this.findRoute(p7ctx.P7request.Method, p7ctx.P7request.URL.Path)
	if nil == p7ri || nil == p7ri.p7node || nil == p7ri.p7node.f4handler {
		p7ctx.I9writer.WriteHeader(404)
		p7ctx.I9writer.Write([]byte(fmt.Sprintf("Not Found:%s %s\r\n", p7ctx.P7request.Method, p7ctx.P7request.URL.Path)))
		return
	}

	p7ctx.M3pathParam = p7ri.m3pathParam
	p7ctx.p7routingNode = p7ri.p7node

	p7ri.p7node.f4handler(p7ctx)
}

func (p7this *HTTPService) Start(addr string) error {
	return http.ListenAndServe(addr, p7this)
}
