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

## 其他

### 如何自定义port

run指定port

```
router.Run(":3000") // 启动3000端口
```

环境变量指定port

```
定义了一个 PORT 的环境变量
```

### Jsoniter

使用 [jsoniter](https://github.com/json-iterator/go) 编译 。Gin 使用 `encoding/json` 作为默认的 json 包，但是你可以在编译中使用标签将其修改为 [jsoniter](https://github.com/json-iterator/go)。  
jsoniter被称为是最快的 JSON 解析器

```php
$ go build -tags=jsoniter .
```

### 运行模式GIN_MODE

目前Gin有三种模式: debug release test 三种，可以通过设置 GIN_MODE 这个环境变量来控制。

- debug：调试模式
- release：发布模式
- test：测试场景