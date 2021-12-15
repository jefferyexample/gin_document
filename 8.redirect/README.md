# 重定向

HTTP 重定向很容易。 内部、外部重定向均支持

```go
r.GET("/baidu", func(c *gin.Context) {
    c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
})
```

路由重定向，使用 HandleContext：

```go
r.GET("/test", func(c *gin.Context) {
    c.Request.URL.Path = "/test2"
    r.HandleContext(c)
})
r.GET("/test2", func(c *gin.Context) {
    c.String(http.StatusOK,"test -> test2 Success")
})
```