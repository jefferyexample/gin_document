# Multipart/Urlencoded

若要将请求主体绑定到结构体中，请使用模型绑定，目前支持JSON、XML、YAML和标准表单值(foo=bar&boo=baz)的绑定。  
Gin使用 [go-playground/validator.v10](https://github.com/go-playground/validator) 验证参数，[查看完整文档](https://pkg.go.dev/gopkg.in/go-playground/validator.v10) 。  
需要在绑定的字段上设置tag，比如，绑定格式为json，需要这样设置 `json:"fieldname"` 。 

## Gin提供的两套绑定方法

### Must bind

Methods - Bind, BindJSON, BindXML, BindQuery, BindYAML  
Behavior - 这些方法底层使用 MustBindWith，如果存在绑定错误，请求将被以下指令中止 `c.AbortWithError(400, err).SetType(ErrorTypeBind)`，响应状态代码会被设置为`400`，请求头`Content-Type`被设置为`text/plain`; `charset=utf-8`。注意，如果你试图在此之后设置响应代码，将会发出一个警告 `[GIN-debug] [WARNING] Headers were already written. Wanted to override status code 400 with 422`，如果你希望更好地控制行为，请使用`ShouldBind`相关的方法

### Should bind

Methods - ShouldBind, ShouldBindJSON, ShouldBindXML, ShouldBindQuery, ShouldBindYAML  
Behavior - 这些方法底层使用 `ShouldBindWith`，如果存在绑定错误，则返回错误，开发人员可以正确处理请求和错误。

### Tips

当我们使用绑定方法时，Gin会根据`Content-Type`推断出使用哪种绑定器，如果你确定你绑定的是什么，你可以使用`MustBindWith`或者`BindingWith`。  
你还可以给字段指定特定规则的修饰符，如果一个字段用`binding:"required"`修饰，并且在绑定时该字段的值为空，那么将返回一个错误。 

## 绑定

```go
type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {
		// 你可以使用显式绑定声明绑定 multipart form：
		// c.ShouldBindWith(&form, binding.Form)
		// 或者简单地使用 ShouldBind 方法自动绑定：
		var form LoginForm
		// 在这种情况下，将自动选择合适的绑定
		if c.ShouldBind(&form) == nil {
			if form.User == "user" && form.Password == "password" {
				c.JSON(200, gin.H{"status": "login success"})
			} else {
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		} else {
			c.JSON(401, gin.H{"status": "missing params"})
		}
	})
	router.Run(":8080")
}
```

测试

```bash
curl -v --form user=user --form password=password http://localhost:8080/login
```

## 表单

```go
func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Multipart/Urlencoded 表单",
		})
	})

	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	router.Run(":8080")
}
```
