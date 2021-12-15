package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// 写法1
	router.LoadHTMLGlob("templates/*")
	router.GET("/1", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "写法1 LoadHTMLGlob",
		})
	})

	// 写法2
	router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/2", func(c *gin.Context) {
		c.HTML(http.StatusOK, "template1.html", gin.H{
			"title": "写法2 LoadHTMLFiles",
		})
	})

	router.Run(":8080")
}
