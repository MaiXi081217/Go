package main

import (
	_ "fmt"
	_ "html/template"
	"net/http"
	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()

	r.Static("/xxx", "./static") //设置静态文件加载目录
	r.LoadHTMLGlob("templates/*")  //设置模板文件目录

	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.Run(":8080")
}
