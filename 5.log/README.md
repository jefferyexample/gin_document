# 日志

## 日志的输出

```go
func main() {
	// 禁用控制台颜色
	gin.DisableConsoleColor()

	// 如果需要控制台输出带有颜色的字体，请使用下面代码
	// gin.ForceConsoleColor()

	// 写入日志文件定义
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	//如果需要将日志输出到控制台，请使用以下代码
	//gin.DefaultWriter = io.MultiWriter(os.Stdout)

	// 如果需要同时将日志写入文件和控制台，请使用以下代码
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "success log")
	})
	router.Run(":8080")
}
```

### 自定义日志格式

