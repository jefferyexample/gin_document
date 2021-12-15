# Cookie和Session

## Cookie

```go
func main() {
	router := gin.Default()

	router.GET("/get_cookie", func(c *gin.Context) {
		cookie, _ := c.Cookie("gin_cookie")
		c.JSON(http.StatusOK, gin.H{
			"get_cookie_text" : cookie,
		})

	})

	router.GET("/set_cookie", func(c *gin.Context) {
		c.SetCookie("gin_cookie", "test Cookie Text", 3600, "/", "*", false, true)
		c.JSON(http.StatusOK, gin.H{
			"message" : "set cookie success",
		})
	})

	router.Run()
}
```

## Session

保存在服务端的键值对数据。  
Session的存在必须依赖于Cookie,Cookie中保存了每个用户Session的唯一标识。  

**引入的包**

```go
"github.com/gin-contrib/sessions"
```

**代码示例**

```go
func main() {
    router := gin.Default()

    // session
    // 创建基于cookie的存储引擎(secretkey 参数是用于加密的密钥)
    store := cookie.NewStore([]byte("secretkey"))
    // 设置session中间件(参数mysession，指的是session的名字，也是cookie的名字)(store是前面创建的存储引擎，我们可以替换成其他存储引擎)
    router.Use(sessions.Sessions("mysession", store))
    router.GET("/get_session", func(c *gin.Context) {
        session := sessions.Default(c) // 初始化session对象
        s := session.Get("test_session") // 设置session数据
        c.JSON(http.StatusOK, gin.H{
            "get_session_text" : s,
        })
    })
    router.GET("/set_session", func(c *gin.Context) {
        session := sessions.Default(c)
        session.Set("test_session", "this is session text")
        session.Save()
        c.JSON(http.StatusOK, gin.H{
            "message" : "set session success",
        })
    })

    router.Run()
}
```
