package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type msg struct {
	// 利用tag对结构体字段定制化
	Name    string `json:"name"` // 当json反射操作，更改name
	Age     int
	Message string
}

func main() {
	r := gin.Default()

	r.GET("/Json", func(c *gin.Context) {
		// 利用map构造JSON
		// data := map[string]interface{}{
		// 	"name":"MaiXi",
		// 	"message":"hello",
		// 	"age":18,
		// }
		/*---------------------------------------------------*/
		// 利用自带的gin.H构造JSON
		data := gin.H{
			"name":    "MaiXi",
			"message": "hello",
			"age":     18,
		}
		c.JSON(http.StatusOK, data)
	})

	r.GET("/another_Json", func(c *gin.Context) {
		data := msg{
			"MaiXi",
			18,
			"hello,world",
		}
		c.JSON(http.StatusOK, data)
	})

	r.Run(":8080")
}
