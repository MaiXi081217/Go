package main

import (
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
)

func posts_hello(c *gin.Context) {
	c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
		"title": "posts",
	})
}
func users_hello(c *gin.Context) {
	c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
		"title": "users",
	})
}

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"safe":func ()  {
			return template.HTML
		}
	})
	// r.LoadHTMLFiles("./templates/index.tmpl")  // 模板解析
	r.LoadHTMLGlob("templates/**/*") // 正则表达
	r.GET("/users/index", users_hello)
	r.GET("/posts/index", posts_hello)

	r.Run(":9090")

}
