# BasicAuth中间件

Basic Auth是一种gin框架提供的认证方式，简单的说就是需要你输入用户名和密码才能继续访问。Bath Auth是其中一种认证方式，另一种是OAuth。  
Basic Authentication是一种HTTP访问控制方式，用于限制对网站资源的访问。这种方式不需要Cookie和Session，只需要客户端发起请求的时候，在头部Header中提交用户名和密码就可以。如果没有附加，会弹出一个对话框，要求输入用户名和密码。这种方式实施起来非常简单，适合路由器之类小型系统。但是它不提供信息加密措施，通常都是以明文或者base64编码传输。  

### 示例代码

```go
func main() {
	r := gin.Default()

	// 路由组使用 gin.BasicAuth() 中间件
	// gin.Accounts 是 map[string]string 的一种快捷方式
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	// /admin/secrets 端点
	// 触发 "localhost:8080/admin/secrets
	authorized.GET("/secrets", func(c *gin.Context) {
		// 获取用户，它是由 BasicAuth 中间件设置的
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	r.Run()
}

// 模拟一些私人数据
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}
```

### 测试

BasicAuth 有三种认证方式，分别是 `浏览器`、`URL提供用户名和密码`、`Authorization认证`

#### 浏览器

浏览器输入 `http://127.0.0.1:8080/admin/secrets` 进行访问。你会看到一个提示框，要你输入用户名和密码。用户名输入foo，密码输入bar。然后你就能看到一个关于foo的信息的JSON串了。

#### URL提供用户名和密码

```bash
curl foo:bar@localhost:8080/admin/secrets
```

通过以上的测试，就已经可以看到测试成功后的json结果了，我们通过 `-v` 看看请求头是怎样的

```bash
curl -v foo:bar@localhost:8080/admin/secrets
```

可以获取到如下参数：

```
*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
* Server auth using Basic with user 'foo'
> GET /admin/secrets HTTP/1.1
> Host: localhost:8080
> Authorization: Basic Zm9vOmJhcg==
> User-Agent: curl/7.77.0
> Accept: */*
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Date: Wed, 15 Dec 2021 09:58:57 GMT
< Content-Length: 64
< 
* Connection #0 to host localhost left intact
{"secret":{"email":"foo@bar.com","phone":"123433"},"user":"foo"}
```

`>`符号后面的是请求，`<`符号后面的是响应。我们注意到请求头中有这样一行请求首部：

```
Authorization: Basic Zm9vOmJhcg==
```

#### Authorization认证

通过上吗得到的Authorization我们进行如下测试

```bash
curl -H 'Authorization: Basic Zm9vOmJhcg==' foo:bar@localhost:8080/admin/secrets
```
