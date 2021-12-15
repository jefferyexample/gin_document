package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type formA struct {
	Foo string `json:"foo" xml:"foo" form:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" form:"bar" binding:"required"`
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.GET("/", HtmlView)
	r.POST("/form1", ShouldBind)
	r.POST("/form2", ShouldBindBodyWith)

	r.Run()
}

func HtmlView(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "将请求绑定到结构体中",
	})
}

func ShouldBind(c *gin.Context) {
	objA := formA{}
	objB := formB{}
	// c.ShouldBind 使用了 c.Request.Body，不可重用。
	if errA := c.ShouldBind(&objA); errA == nil {
		a, _ := json.Marshal(objA)
		c.String(http.StatusOK, "the body should be formA："+string(a))
		// 因为现在 c.Request.Body 是 EOF，所以这里会报错。
	} else if errB := c.ShouldBind(&objB); errB == nil {
		b, _ := json.Marshal(objB)
		c.String(http.StatusOK, "the body should be formB："+string(b))
	} else {
		c.String(http.StatusOK, "bind error")
	}
}

func ShouldBindBodyWith(c *gin.Context) {
	objA := formA{}
	objB := formB{}
	// 读取 c.Request.Body 并将结果存入上下文。
	if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
		a, _ := json.Marshal(objA)
		c.String(http.StatusOK, "the body should be formA："+string(a))
		// 这时, 复用存储在上下文中的 body。
	} else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
		b, _ := json.Marshal(objB)
		c.String(http.StatusOK, "the body should be formB Json："+string(b))
		// 可以接受其他格式
	} else if errB2 := c.ShouldBindBodyWith(&objB, binding.XML); errB2 == nil {
		c.String(http.StatusOK, `the body should be formB XML`)
	} else {
		c.String(http.StatusOK, "bind error")
	}
}

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
