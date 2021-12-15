# 路由

路由系统的用于接收请求，并将请求转发给注册的中间件或请求处理器来处理。核心功能如下：

- 路由系统可根据请求方法，请求路径，和路径参数来识别转发。
- 可设置一个或多个中间件用于在请求处理器前后，处理特殊的事件。
- 可以分组设置，将一个或多个中间件作用在一组多个路由上。

## 快速开始

开始使用 GET, POST, PUT, PATCH, DELETE and OPTIONS

```go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/get", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This get",
		})
	})

	r.POST("/post", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This post",
		})
	})

	r.PUT("/put", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This put",
		})
	})

	r.Run()
}
```

## 路由参数

### Params

**路由参数**

指的是在路由路径中定义的参数，例如请求 URI 是 /user/42，42 作为用户的 ID，那么 42 就是路由参数。注意路由参数不是查询字符串，例如 /user?ID=42， 这个 ID=42，也就是问号后边的才是查询字符串。  
使用路由参数的好处是将动态 URL 变为 静态 URL，因为请求客户端会认为路由参数不是变化的数据，因此被视为静态 URL。  
若需要定义带有路由参数的路径，需要使用 :param 或 *param 的语法在路径中。  

**格式**

```
# 定义
/router1/:id

# 请求
/router1/123
/router1/789
```

**示例**

```go
func main() {
    r := gin.Default()
    
    // 匹配成功：/user/jack
    // 匹配失败：/user/
    r.GET("/user/:name", func(c *gin.Context) {
        name := c.Param("name")
        c.JSON(200, gin.H{
        "message": "Param：" + name,
        })
    })
    
    // 匹配成功：/need/jack/tt 、 /need/jack/
    // 匹配失败：/need/ 、 /need/jack/
    r.GET("/need/:name/*action", func(c *gin.Context) {
        name := c.Param("name")
        action := c.Param("action")
        path := c.FullPath() // 请求：/need/jack/tt 、 输出：/need/:name/*action
        c.String(http.StatusOK, "Path:" + path + " Params: " + name + " is " + action)
    })
    
    r.Run()
}
```

**必选参数**

使用 :param 的语法完成必选参数，例如 /user/:ID

**可选参数**

使用 *param 的语法完成可选参数的定义，例如 /user/*ID

**获取参数**

使用 gin.Context 对象的 c.Param("param") 来获取参数值，参见上面的示例。

### Query

格式

```
/router1?id=123
/router1?id=456
```

示例

```go
func main() {
    r := gin.Default()
    r.GET("/welcome", func(c *gin.Context) {
        q1 := c.DefaultQuery("q1", "Guest")
        // c.Request.URL.Query().Get("lastname") 的快捷写法
        q2 := c.Query("q2")
        c.JSON(200, gin.H{
            "q1" : q1,
            "q2" : q2,
        })
    })
    r.Run()
}
```

## 接收参数

**获取GET参数**

```go
firstname := c.DefaultQuery("firstname", "Guest")
lastname := c.Query("lastname") // 是 c.Request.URL.Query().Get("lastname") 的简写
```

**获取POST参数**

```go
message := c.PostForm("message")
nick := c.DefaultPostForm("nick", "anonymous") // 此方法可以设置默认值
```

## 路由分组

使用函数 router.Group() 完成分组的创建。创建时可以提供路径前缀和公用中间件，router.Group() 函数签名如下：

```go
func (group *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup
```

调动该函数后，会形成一个分组路由对象，组内的路由需要使用该对象完成处理器的注册，例如：

```go
func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("read", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "This /v1/read",
			})
		})
		v1.GET("login", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "This /v1/login",
			})
		})
	}
	r.Run()
}
```

上面的代码创建了一个路由分组，前缀是 v1 。本例中，并没有设置任何的中间件，这是可行的。  
注意，v1 组的路由的注册方案，都是由 v1. 调用。  
语法上，我们习惯（也是官方建议）将一组路由放在一个代码块中，在结构上保持独立。但这个代码块不是必要的。  