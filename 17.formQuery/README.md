# 映射查询字符串或表单参数

**示例代码**

```go
func main() {
    router := gin.Default()

    router.POST("/post", func(c *gin.Context) {

        ids := c.QueryMap("ids")
        names := c.PostFormMap("names")

        fmt.Printf("ids: %v; names: %v", ids, names)
    })
    router.Run(":8080")
}
```

**测试**

```bash
POST /post?ids[a]=1234&ids[b]=hello HTTP/1.1
Content-Type: application/x-www-form-urlencoded
names[first]=thinkerou&names[second]=tianou
```

返回

```
ids: map[b:hello a:1234], names: map[second:tianou first:thinkerou]
```