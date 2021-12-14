# 中间件

中间件 middleware，也是一种处理器。主要用于在多个业务逻辑中间重用代码，例如认证校验，日志处理等。  
我们在使用 `gin.Default()` 初始化路由对象时，会随之附加两个中间件 `Logger` 和 `Recovery` ，用于完成日志和恢复的相关处理，参考的 `gin.Default()` 代码如下：  

## Gin中的默认中间件

```go
r := gin.Default()

r := gin.New()
```

### New 和 Default 的区别

Default 默认已经连接了 Logger 和 Recovery 中间件

**Logger**

全局中间件，Logger 中间件将写日志到 `gin.DefaultWriter` 即使你设置 `GIN_MODE=release`  
默认设置 `gin.DefaultWriter = os.Stdout`  

**Recovery**

Recovery 中间件从任何 panic 恢复，如果出现 panic，它会写一个 500 错误。

### 用New实现Default的中间件

```go
func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.Run()
}
```

## 使用中间件

### 自定义中间件

我们已经知道，Gin的中间件其实就是一个HandlerFunc,那么只要我们自己实现一个HandlerFunc，就可以自定义一个自己的中间件。接下来演示如何自定义一个中间件。

#### 全局中间件

```go
func main() {
	// 全局中间件
	r := gin.New()

	// 加载中间件的两种方式
	r.Use(middle1)
	r.Use(middle2())

	r.GET("/get", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"message" : "加载两种中间件",
		})
	})

	r.Run()
}

// 自定义中间件
// 方法1
func middle1(c *gin.Context)  {
	fmt.Println("中间件1之前")
	c.Next()
	fmt.Println("中间件1之后...")
}

// 自定义中间件
// 方法2
func middle2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Middle 2 start")
		c.Next()
		fmt.Println("Middle 2 End...")
	}
}
```

**c.Next方法**

这个是执行后续中间件请求处理的意思。  
在当前中间件中调用`c.Next()`时会中断当前中间件中后续的逻辑，转而执行后续的中间件和`handlers`，等他们全部执行完以后再回来执行当前中间件的后续代码。

**c.Next使用总结**

- 中间件代码最后即使没有调用Next()方法，后续中间件及handlers也会执行；
- 如果在中间件函数的非结尾调用Next()方法当前中间件剩余代码会被暂停执行，会先去执行后续中间件及handlers，等这些handlers全部执行完以后程序控制权会回到当前中间件继续执行剩余代码；
- 如果想提前中止当前中间件的执行应该使用return退出而不是Next()方法；
- 如果想中断剩余中间件及handlers应该使用Abort方法，但需要注意当前中间件的剩余代码会继续执行。

#### 局部中间件

```go
func main() {
	r := gin.New()

	// 在路由中添加局部中间件
	r.GET("/get", middle1, func(c *gin.Context) {
		c.JSON(200,gin.H{
			"message" : "局部中间件 /get",
		})
	})

	// 在路由组中添加局部中间件
	v1 := r.Group("/v1")
	v1.Use(middle1)
	{
		v1.GET("read", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "This middle /v1/read",
			})
		})
	}

	r.Run()
}
```