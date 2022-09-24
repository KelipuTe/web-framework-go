package v40

import (
	"fmt"
	"net/http"
)

// HTTPContext 自定义请求上下文（注意和 context.Context 的概念区分开）
type HTTPContext struct {
	// ServeHTTP 的 http.ResponseWriter
	I9writer http.ResponseWriter
	// ServeHTTP 的 *http.Request
	P7request *http.Request

	// 命中的路由结点
	p7routingNode *routingNode
	// 提取到的路径参数
	M3pathParam map[string]string
}

// GetRoutingInfo 获取命中的路由结点的信息
func (this HTTPContext) GetRoutingInfo() string {
	return fmt.Sprintf("nodeType:%d\r\nrouting path:%s\r\n", this.p7routingNode.nodeType, this.p7routingNode.path)
}
