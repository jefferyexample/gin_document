package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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
		c.String(http.StatusOK, "Path:"+path+" Params: "+name+" is "+action)
	})

	// 匹配成功：/welcome?q1=tt&q2=666
	r.GET("/welcome", func(c *gin.Context) {
		q1 := c.DefaultQuery("q1", "Guest")
		// c.Request.URL.Query().Get("lastname") 的快捷写法
		q2 := c.Query("q2")
		c.JSON(200, gin.H{
			"q1": q1,
			"q2": q2,
		})
	})

	// 路由分组
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
