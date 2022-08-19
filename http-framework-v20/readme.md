## 说明（readme）

http_service_v1 是一个简单的 HTTP 框架。

为方便理解框架演进的过程 http_service_v1 作为的示例代码就不动了。

拷贝一份代码出来到 http_service_v2，后续演进在 http_service_v2 上继续。

简单的实现了：

- 中间件
- 基于 map 的路由（静态匹配）
- 基于前缀树的路由 V1（静态匹配，模糊匹配）