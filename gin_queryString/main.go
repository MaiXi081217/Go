package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func main(){
	r := gin.Default()

	//处理 get请求在url后面?后面的querystring参数
	//key=value格式，多个key-value用&连接
	//eq: /web?query=xxx&age=xxx
	r.GET("/web",func(c *gin.Context) {
		//获取浏览器端发送的querystring参数
		/*------------------------------------------------------------------*/
		// 设定somebody为默认值
		//name := c.DefaultQuery("query","somebody")
		/*------------------------------------------------------------------*/
		// 取不到query返回bool
		// 	name,ok := c.GetQuery("query") 
		// 	if !ok{
		// 		name = "somebody"
		// 	}
		name := c.Query("query")
		age := c.Query("age")
		c.JSON(http.StatusOK,gin.H{
				"name":name,
				"age":age,
			})
		})
		/*------------------------------------------------------------------*/
		//传送多个query键值对参数


	r.Run(":9090")
}