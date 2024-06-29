package main

import (
	"net/http"
	"html/template"

	"github.com/gin-gonic/gin"
)

func posts_hello(c *gin.Context) {
	c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
		"title": "posts",
	})
}
func users_hello(c *gin.Context) {
	c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
		"title": "<a href = 'https://www.bilibili.com/'>mc的github</a>",
	})
}

func main() {
	r := gin.Default()
	// 给模板添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe":func(str string) template.HTML{
			return template.HTML(str)
		},
	}) 
	// r.LoadHTMLFiles("./templates/index.tmpl")  // 模板解析
	r.LoadHTMLGlob("templates/**/*") // 正则表达
	r.GET("/users/index", users_hello)
	r.GET("/posts/index", posts_hello)

	r.Run(":9090")

}
