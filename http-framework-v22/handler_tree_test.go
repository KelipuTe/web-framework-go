package http_framework_v22

import (
  "errors"
  "net/http"
  "reflect"
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestHTTPHandlerTree_RegisteRoute(t *testing.T) {
  // 创建接口实例，记得断言成测试的对象
  p1h := NewHTTPHandlerTree().(*HTTPHandlerTree)
  // 检查支持的 HTTP 方法数量是否正确
  assert.Equal(t, len(arr1supportedMethod), len(p1h.mapp1root))

  // 添加 PUT /user
  err := p1h.RegisteRoute(http.MethodPut, "/user", func(p1c *HTTPContext) {})
  // 检查是否抛出异常
  assert.Equal(t, errors.New("method not supported"), err)

  // 添加 GET /user
  err = p1h.RegisteRoute(http.MethodGet, "/user", func(p1c *HTTPContext) {})
  assert.Nil(t, err)

  // 检查 GET 组 / 结点，子结点数量
  t1getNode := p1h.mapp1root[http.MethodGet]
  assert.Equal(t, 1, len(t1getNode.arr1p1children))
  // 检查 GET 组 /user 结点
  t1userNode := t1getNode.arr1p1children[0]
  assert.NotNil(t, t1userNode)
  assert.Equal(t, "user", t1userNode.pattern)
  assert.Empty(t, t1userNode.arr1p1children)
  assert.NotNil(t, t1userNode.hhFunc)

  // 添加 GET /user/info
  err = p1h.RegisteRoute(http.MethodGet, "/user/info", func(p1c *HTTPContext) {})
  assert.Nil(t, err)
  // 检查 GET 组 /user 结点，子结点数量
  assert.Equal(t, 1, len(t1userNode.arr1p1children))
  // 检查 GET 组 /user/info 结点
  t1userInfoNode := t1userNode.arr1p1children[0]
  assert.NotNil(t, t1userInfoNode)
  assert.Equal(t, "info", t1userInfoNode.pattern)
  assert.Empty(t, t1userInfoNode.arr1p1children)
  assert.NotNil(t, t1userInfoNode.hhFunc)

  // 重复添加 GET /user/info
  err = p1h.RegisteRoute(http.MethodGet, "/user/info", func(p1c *HTTPContext) {})
  assert.Nil(t, err)
  // 检查 GET 组 /user 结点，子结点数量
  assert.Equal(t, 1, len(t1userNode.arr1p1children))
  // 检查 GET 组 /user/info 结点
  assert.NotNil(t, t1userInfoNode)
  assert.Equal(t, "info", t1userInfoNode.pattern)
  assert.Empty(t, t1userInfoNode.arr1p1children)
  assert.NotNil(t, t1userInfoNode.hhFunc)

  // 添加 GET /user/detail
  err = p1h.RegisteRoute(http.MethodGet, "/user/detail", func(p1c *HTTPContext) {})
  assert.Nil(t, err)
  // 检查 GET 组 /user 结点，子结点数量
  assert.Equal(t, 2, len(t1userNode.arr1p1children))
  t1userDetailNode := t1userNode.arr1p1children[1]
  assert.NotNil(t, t1userDetailNode)
  assert.Equal(t, "detail", t1userDetailNode.pattern)
  assert.Empty(t, t1userDetailNode.arr1p1children)
  assert.NotNil(t, t1userDetailNode.hhFunc)

  // 添加 POST /order/check
  err = p1h.RegisteRoute(http.MethodPost, "/order/check", func(p1c *HTTPContext) {})
  assert.Nil(t, err)

  // 检查 POST 组 / 结点，子结点数量
  t1postNode := p1h.mapp1root[http.MethodPost]
  assert.Equal(t, 1, len(t1postNode.arr1p1children))
  // 检查 POST 组 /order 结点
  t1orderNode := t1postNode.arr1p1children[0]
  assert.NotNil(t, t1orderNode)
  assert.Equal(t, "order", t1orderNode.pattern)
  assert.Equal(t, 1, len(t1orderNode.arr1p1children))
  assert.Nil(t, t1orderNode.hhFunc)
  // 检查 POST 组 /order/check 结点
  t1orderCheckNode := t1orderNode.arr1p1children[0]
  assert.NotNil(t, t1orderCheckNode)
  assert.Equal(t, "check", t1orderCheckNode.pattern)
  assert.Empty(t, t1orderCheckNode.arr1p1children)
  assert.NotNil(t, t1orderCheckNode.hhFunc)

  // 添加 POST /order/*
  err = p1h.RegisteRoute(http.MethodPost, "/order/*", func(p1c *HTTPContext) {})
  assert.Nil(t, err)
  // 检查 POST 组 /order 结点
  assert.Equal(t, 2, len(t1orderNode.arr1p1children))
  // 检查 POST 组 /order/* 结点
  t1orderAnyNode := t1orderNode.arr1p1children[1]
  assert.NotNil(t, t1orderAnyNode)
  assert.Equal(t, "*", t1orderAnyNode.pattern)
  assert.Empty(t, t1orderAnyNode.arr1p1children)
  assert.NotNil(t, t1orderAnyNode.hhFunc)

  // 添加 POST *
  err = p1h.RegisteRoute(http.MethodPost, "*", func(p1c *HTTPContext) {})
  // 检查是否抛出异常
  assert.Equal(t, errors.New("route pattern is error, index == 0"), err)
  // 添加 POST /order*
  err = p1h.RegisteRoute(http.MethodPost, "/order*", func(p1c *HTTPContext) {})
  // 检查是否抛出异常
  assert.Equal(t, errors.New("route pattern is error, '/' != pattern[index-1]"), err)
  // 添加 POST /order/*/detail
  err = p1h.RegisteRoute(http.MethodPost, "/order/*/detail", func(p1c *HTTPContext) {})
  // 检查是否抛出异常
  assert.Equal(t, errors.New("route pattern is error, len(pattern) - 1 != index"), err)

  // 添加 POST /order/:id
  err = p1h.RegisteRoute(http.MethodPost, "/order/:id", func(p1c *HTTPContext) {})
  assert.Nil(t, err)
  assert.Equal(t, 3, len(t1orderNode.arr1p1children))
  t1orderIdNode := t1orderNode.arr1p1children[2]
  assert.NotNil(t, t1orderAnyNode)
  assert.Equal(t, ":id", t1orderIdNode.pattern)
  assert.Empty(t, t1orderAnyNode.arr1p1children)
  assert.NotNil(t, t1orderAnyNode.hhFunc)
}

func TestHTTPHandlerTree_(t *testing.T) {
  // 创建接口实例，记得断言成测试的对象
  p1h := NewHTTPHandlerTree().(*HTTPHandlerTree)
  p1c := NewHTTPContext()

  // 添加 GET /user
  _ = p1h.RegisteRoute(http.MethodGet, "/user", func(p1c *HTTPContext) {})

  hhFunc, err := p1h.findRoute(http.MethodGet, "/user", p1c)
  assert.Nil(t, err)
  assert.NotNil(t, hhFunc)

  hhFunc, err = p1h.findRoute(http.MethodGet, "/user/info", p1c)
  assert.Equal(t, errors.New("route not found"), err)

  // 添加 GET /user/info
  _ = p1h.RegisteRoute(http.MethodGet, "/user/info", func(p1c *HTTPContext) {})

  hhFunc, err = p1h.findRoute(http.MethodGet, "/user/info", p1c)
  assert.Nil(t, err)
  assert.NotNil(t, hhFunc)

  // 添加 POST /order/check
  _ = p1h.RegisteRoute(http.MethodPost, "/order/check", func(p1c *HTTPContext) {})

  hhFunc, err = p1h.findRoute(http.MethodPost, "/order", p1c)
  assert.NotNil(t, err)

  hhFunc, err = p1h.findRoute(http.MethodPost, "/order/check", p1c)
  assert.Nil(t, err)
  assert.NotNil(t, hhFunc)

  // 添加 POST /order/*
  var f1orderAny HTTPHandlerFunc = func(p1c *HTTPContext) {}
  _ = p1h.RegisteRoute(http.MethodPost, "/order/*", f1orderAny)

  hhFunc, err = p1h.findRoute(http.MethodPost, "/order/any", p1c)
  assert.Nil(t, err)
  assert.NotNil(t, hhFunc)
  assert.True(t, hhFuncCheckEqual(hhFunc, f1orderAny))

  // 添加 POST /order/:id
  var f1orderId HTTPHandlerFunc = func(p1c *HTTPContext) {}
  _ = p1h.RegisteRoute(http.MethodPost, "/order/:id", f1orderId)

  hhFunc, err = p1h.findRoute(http.MethodPost, "/order/123", p1c)
  assert.Nil(t, err)
  assert.NotNil(t, hhFunc)
  assert.True(t, hhFuncCheckEqual(hhFunc, f1orderId))

  // 添加 POST /order/:id/detail
  var f1orderIdDetail HTTPHandlerFunc = func(p1c *HTTPContext) {}
  _ = p1h.RegisteRoute(http.MethodPost, "/order/:id/detail", f1orderIdDetail)

  hhFunc, err = p1h.findRoute(http.MethodPost, "/order/123/detail", p1c)
  assert.Nil(t, err)
  assert.NotNil(t, hhFunc)
  assert.True(t, hhFuncCheckEqual(hhFunc, f1orderIdDetail))
}

// 检查两个方法是否是同一个
func hhFuncCheckEqual(hhFuncA, hhFuncB HTTPHandlerFunc) bool {
  return reflect.ValueOf(hhFuncA).Pointer() == reflect.ValueOf(hhFuncB).Pointer()
}
