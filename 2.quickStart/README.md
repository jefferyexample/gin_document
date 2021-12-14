# 快速开始

```go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Gin",
		})
	})
	r.Run()
}
```

在终端运行以下命令

```bash
go run main.go
```

在浏览器访问

```
http://127.0.0.1:8080/
```

## 如何自定义port

run指定port

```
router.Run(":3000") // 启动3000端口
```

环境变量指定port

```
定义了一个 PORT 的环境变量
```