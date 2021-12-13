# Gin 框架介绍

Gin是用Go（Golang）编写的Web框架。他是一个类似于 [martini](https://github.com/go-martini/martini) 但拥有更好性能的API框架，由于 [httprouter](https://github.com/julienschmidt/httprouter)，速度提高了40倍。如果您追求性能和高效的效率，您将会爱上Gin。

## 安装

在安装Gin包之前，你需要在你的电脑上安装Go环境并设置你的工作区。

1. 首先需要安装Go(支持版本1.11+)，然后使用以下Go命令安装Gin:

```bash
go get -u github.com/gin-gonic/gin
```
   
2. 在你的代码中导入Gin包: 

```bash
import "github.com/gin-gonic/gin" 
```

3. (可选)如果使用诸如 http.StatusOK 之类的常量，则需要引入 net/http 包：

```bash
import "net/http"
```