# 静态文件服务

如果你想让gin实现一个类似于nginx的服务，为你的静态文件提供和一个server服务，可以用以下方法实现。

## 代码示例

```go
func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	// 监听并在 0.0.0.0:8080 上启动服务
	router.Run(":8080")
}
```

**参数说明：**

- router.Static 指定某个目录为静态资源目录，可直接访问这个目录下的资源，url 要具体到资源名称。
- router.StaticFS 比前面一个多了个功能，当目录下不存 index.html 文件时，会列出该目录下的所有文件。
- router.StaticFile 指定某个具体的文件作为静态资源访问。