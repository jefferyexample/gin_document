package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

//func main() {
//	// 禁用控制台颜色
//	gin.DisableConsoleColor()
//
//	// 如果需要控制台输出带有颜色的字体，请使用下面代码
//	// gin.ForceConsoleColor()
//
//	// 写入日志文件定义
//	f, _ := os.Create("gin.log")
//	gin.DefaultWriter = io.MultiWriter(f)
//
//	//如果需要将日志输出到控制台，请使用以下代码
//	//gin.DefaultWriter = io.MultiWriter(os.Stdout)
//
//	// 如果需要同时将日志写入文件和控制台，请使用以下代码
//	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
//
//	router := gin.Default()
//	router.GET("/", func(c *gin.Context) {
//		c.String(200, "success log")
//	})
//	router.Run(":8080")
//}

func main() {
	r := gin.Default()

	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Dyi log fmt")
	})
	r.Run(":8080")
}
