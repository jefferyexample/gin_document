# 模版渲染

使用 LoadHTMLGlob () 或者 LoadHTMLFiles () 进行模版渲染

```go
func main() {
	router := gin.Default()

	// 写法1
	router.LoadHTMLGlob("templates/*")
	router.GET("/1", func(c *gin.Context) {
		c.HTML(http.StatusOK, "template1.html", gin.H{
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
```

**使用不同目录下名称相同的模板**

```go
func main() {
    router := gin.Default()
    router.LoadHTMLGlob("templates/**/*")
    router.GET("/posts/index", func(c *gin.Context) {
        c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
            "title": "Posts",
        })
    })
    router.GET("/users/index", func(c *gin.Context) {
        c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
            "title": "Users",
        })
    })
    router.Run(":8080")
}
```

templates/posts/index.tmpl

```go
{{ define "posts/index.tmpl" }}
<html><h1>
    {{ .title }}
</h1>
<p>Using posts/index.tmpl</p>
</html>
{{ end }}
```

templates/users/index.tmpl

```go
{{ define "users/index.tmpl" }}
<html><h1>
    {{ .title }}
</h1>
<p>Using users/index.tmpl</p>
</html>
{{ end }}
```

## 自定义模板渲染器

你可以使用自定义的 html 模板渲染

```go
import "html/template"

func main() {
    router := gin.Default()
    html := template.Must(template.ParseFiles("file1", "file2"))
    router.SetHTMLTemplate(html)
    router.Run(":8080")
}
```

自定义分隔符

```go
r := gin.Default()
r.Delims("{[{", "}]}")
r.LoadHTMLGlob("/path/to/templates")
```

自定义模板功能

```go
import (
    "fmt"
    "html/template"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

func formatAsDate(t time.Time) string {
    year, month, day := t.Date()
    return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func main() {
    router := gin.Default()
    router.Delims("{[{", "}]}")
    router.SetFuncMap(template.FuncMap{
        "formatAsDate": formatAsDate,
    })
    router.LoadHTMLFiles("./testdata/template/raw.tmpl")

    router.GET("/raw", func(c *gin.Context) {
        c.HTML(http.StatusOK, "raw.tmpl", map[string]interface{}{
            "now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
        })
    })

    router.Run(":8080")
}
```

raw.tmpl

```gotemplate
Date: {[{.now | formatAsDate}]}
```

结果：

```
Date: 2017/07/01
```

## 多模版

Gin 默认允许只使用一个 html 模板。 查看 [多模板渲染](https://github.com/gin-contrib/multitemplate) 以使用 go 1.6 block template 等功能。