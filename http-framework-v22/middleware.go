package http_framework_v22

import (
  "fmt"
  "time"
)

// MiddlewareFunc 中间件处理方法。
type MiddlewareFunc func(c *HTTPContext)

// MiddlewareBuilder 中间件建造器。
type MiddlewareBuilder func(next MiddlewareFunc) MiddlewareFunc

// Test1MiddlewareBuilder 测试调用顺序
func TestMiddlewareBuilder(next MiddlewareFunc) MiddlewareFunc {
  return func(c *HTTPContext) {
    fmt.Printf("request before test1 middleware.\n")
    next(c)
    fmt.Printf("request after test1 middleware.\n")
  }
}

// TimeCostMiddlewareBuilder 算一下耗时
func TimeCostMiddlewareBuilder(next MiddlewareFunc) MiddlewareFunc {
  return func(c *HTTPContext) {
    startUN := time.Now().UnixNano()
    next(c)
    endUN := time.Now().UnixNano()
    fmt.Printf("request time cost: %d unix nano.\n", startUN-endUN)
  }
}
