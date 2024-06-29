package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func posts_hello(c *gin.Context) {
	c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
		"title": "posts",
	})
}
func users_hello(c *gin.Context) {
	c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
		"title": "<a href = 'https://www.bilibili.com/'>Tis is bilibili</a>",
	})
}

func String_To_Html(str string) template.HTML {
	return template.HTML(str)
}

func main() {
	r := gin.Default()
	// 加载静态文件
	r.Static("/xxx", "./static")
	// 给模板添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": String_To_Html,
	})
	// r.LoadHTMLFiles("./templates/index.tmpl")  // 模板解析
	r.LoadHTMLGlob("templates/**/*") // 正则表达
	r.GET("/users/index", users_hello)
	r.GET("/posts/index", posts_hello)
	r.GET("/posts/home",func(c *gin.Context) {
		c.HTML(http.StatusOK,"/home.html",nil)
	})
	r.Run(":9090")

}
