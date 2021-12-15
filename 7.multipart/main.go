package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

//func main() {
//	router := gin.Default()
//	router.POST("/login", func(c *gin.Context) {
//		// 你可以使用显式绑定声明绑定 multipart form：
//		// c.ShouldBindWith(&form, binding.Form)
//		// 或者简单地使用 ShouldBind 方法自动绑定：
//		var form LoginForm
//		// 在这种情况下，将自动选择合适的绑定
//		if c.ShouldBind(&form) == nil {
//			if form.User == "user" && form.Password == "password" {
//				c.JSON(200, gin.H{"status": "login success"})
//			} else {
//				c.JSON(401, gin.H{"status": "unauthorized"})
//			}
//		} else {
//			c.JSON(401, gin.H{"status": "missing params"})
//		}
//	})
//	router.Run(":8080")
//}

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
