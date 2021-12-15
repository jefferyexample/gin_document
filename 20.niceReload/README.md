# 优雅重启或停止

想要优雅地重启或停止你的Web服务器，使用下面的方法

我们可以使用[fvbock/endless](https://github.com/fvbock/endless)来替换默认的`ListenAndServe`，有关详细信息，请参阅问题[＃296](https://github.com/gin-gonic/gin/issues/296)

```
router := gin.Default()
router.GET("/", handler)
// [...]
endless.ListenAndServe(":4242", router)
```

一个替换方案

- [manners](https://github.com/braintree/manners) ：一个Go HTTP服务器，能优雅的关闭
- [graceful](https://github.com/tylerb/graceful) ：Graceful是一个go的包，支持优雅地关闭http.Handler服务器
- [grace](https://github.com/facebookgo/grace) ：对Go服务器进行优雅的重启和零停机部署

如果你的Go版本是1.8，你可能不需要使用这个库，考虑使用http.Server内置的 [Shutdown()](https://golang.org/pkg/net/http/#Server.Shutdown) 方法进行优雅关闭，查看 [例子](https://github.com/gin-gonic/gin/tree/master/examples/graceful-shutdown)

```go
// +build go1.8

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
```