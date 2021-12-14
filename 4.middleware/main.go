package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 全局中间件
//func main() {
//	r := gin.New()
//	r.Use(gin.Logger(), gin.Recovery())
//	r.Run()
//}

// 全局中间件
//func main() {
//	r := gin.New()
//
//	// 加载中间件的两种方式
//	r.Use(middle1)
//	r.Use(middle2())
//
//	r.GET("/get", func(c *gin.Context) {
//		c.JSON(200,gin.H{
//			"message" : "加载两种中间件",
//		})
//	})
//
//	r.Run()
//}

// 自定义中间件
// 方法1
func middle1(c *gin.Context) {
	fmt.Println("中间件1之前")
	c.Next()
	fmt.Println("中间件1之后...")
}

// 自定义中间件
// 方法2
//func middle2() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		fmt.Println("Middle 2 start")
//		c.Next()
//		fmt.Println("Middle 2 End...")
//	}
//}

// 局部中间件
func main() {
	r := gin.New()

	// 在路由中添加局部中间件
	r.GET("/get", middle1, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "局部中间件 /get",
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
