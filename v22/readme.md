## 说明（readme）

http_service_v2 是一个简单的 HTTP 框架。基于 http_service_v1 继续演进。

http_service_v2 只会对演进的部分加注释，http_service_v1 写过的注释 http_service_v2 不会重复一遍。

简单的实现了：

- 中间件
- 基于前缀树的路由 V2（静态匹配，模糊匹配，路由参数）
- 资源复用（主要是上下文资源 HTTPContext）
- 服务关闭（信号处理、拒接新请求、等待正在运行的请求）
- 单元测试
