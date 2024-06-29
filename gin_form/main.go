package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./index.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.POST("/login", func(c *gin.Context) {
		// username := c.PostForm("username") // 取的是输入对象的name
		// password := c.PostForm("password")
		// username := c.DefaultPostForm("username", "123")
		// password := c.DefaultPostForm("xx", "123")  //如果属性名在页面未定义，返回默认值
		username, ok := c.GetPostForm("username") //如果取不到数值，返回默认值
		if !ok {
			username = "MaiXi"
		}
		password, ok := c.GetPostForm("password")
		if !ok {
			password = "NULL"
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"Password": password,
		})
	})
	r.Run(":9090")
}
