# 将请求绑定到结构体中

Gin 支持对不同传参方式的参数进行统一绑定并验证

## 常见的绑定方式

context.Bind() 都可以绑定  
context.ShouldBind() 都可以绑定  
ShouldBindQuery() 只能绑定get  

## Form

绑定请求体的常规方法使用c.Request.Body，并且不能多次调用

**支持格式**

```
Content-Type: application/x-www-form-urlencoded with a=XX&b=0
```

**代码示例**

```go
type formA struct {
Foo string `json:"foo" xml:"foo" form:"foo" binding:"required"`
}

type formB struct {
Bar string `json:"bar" xml:"bar" form:"bar" binding:"required"`
}

func ShouldBind(c *gin.Context)  {
	objA := formA{}
	objB := formB{}
	// c.ShouldBind 使用了 c.Request.Body，不可重用。
	if errA := c.ShouldBind(&objA); errA == nil {
		a, _ := json.Marshal(objA)
		c.String(http.StatusOK, "the body should be formA：" + string(a))
		// 因为现在 c.Request.Body 是 EOF，所以这里会报错。
	} else if errB := c.ShouldBind(&objB); errB == nil {
		b, _ := json.Marshal(objB)
		c.String(http.StatusOK, "the body should be formB：" + string(b))
	} else {
		c.String(http.StatusOK, "bind error")
	}
}
```

## Json

同样，你能使用c.ShouldBindBodyWith  
c.ShouldBindBodyWith 在绑定之前将body存储到上下文中，这对性能有轻微影响，因此如果你要立即调用，则不应使用此方法  
此功能仅适用于这些格式 -- JSON, XML, MsgPack, ProtoBuf。对于其他格式，Query, Form, FormPost, FormMultipart, 可以被c.ShouldBind()多次调用而不影响性能（参考 [#1341](https://github.com/gin-gonic/gin/pull/1341)）  

**支持格式**

```
Content-Type: application/json with { "a":"XX", "b":0 }
```

**代码示例**

```go
type formA struct {
Foo string `json:"foo" xml:"foo" form:"foo" binding:"required"`
}

type formB struct {
Bar string `json:"bar" xml:"bar" form:"bar" binding:"required"`
}

func ShouldBindBodyWith(c *gin.Context)  {
	objA := formA{}
	objB := formB{}
	// 读取 c.Request.Body 并将结果存入上下文。
	if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
		a, _ := json.Marshal(objA)
		c.String(http.StatusOK, "the body should be formA：" + string(a))
		// 这时, 复用存储在上下文中的 body。
	} else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
		b, _ := json.Marshal(objB)
		c.String(http.StatusOK, "the body should be formB Json：" + string(b))
		// 可以接受其他格式
	} else if errB2 := c.ShouldBindBodyWith(&objB, binding.XML); errB2 == nil {
		c.String(http.StatusOK, `the body should be formB XML`)
	} else {
		c.String(http.StatusOK, "bind error")
	}
}
```

**测试**

```bash
curl 'http://127.0.0.1:8080/form2' -H 'Content-Type: application/json' -d '{"foo":"111", "bar":"222"}'
```

## 总结

### ShouldBind多次绑定错误问题

- 参考：https://github.com/gin-gonic/gin/pull/1341

`ShouldBind` 方法是最常用解析JSON数据的方法之一，但在重复调用的情况下会出现EOF的报错，这个原因出在`ShouldBind`在调用过一次之后`context.request.body.sawEOF`的值是`false`导致，所以如果要多次绑定多个变量，需要使用`ShouldBindBodyWith`。  

以下为使用`单次绑定`和`多次绑定`的代码演示示例：

```go
// MsgJson Json数据结构体
type MsgJson struct {
	Msg string `json:"msg"`
}

// 单次绑定请使用 ShouldBindJSON 方法
func bindExample(c *gin.Context) {
	var a MsgJson
	if err := c.ShouldBindJSON(&a); err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
	return
}

// 多次绑定请使用 ShouldBindBodyWith 方法
func bindWithRightWay(c *gin.Context) {
	var a, b MsgJson

	// 第一次绑定与解析
	if err := c.ShouldBindBodyWith(&a, binding.JSON); err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	// 第二次绑定与解析
	if err := c.ShouldBindBodyWith(&b, binding.JSON); err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	// 返回
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
	return
}
```

### EOF错误复现

如果对 `ShouldBind` 进行多次绑定，则会出现 `EOF` 错误。以下代码为错误示例：

```go
func bindWithError(c *gin.Context) {
	var a, b MsgJson
	if err := c.ShouldBindJSON(&a); err != nil {....}

	// ---> 注意，这里会出现EOF报错
	if err := c.ShouldBindJSON(&b); err != nil {....}
	.......
	return
}
```

