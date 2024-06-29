package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct{
	username string
	password string
}

func main(){
	r := gin.Default()
	r.GET("/user",func(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password")
		u := UserInfo{
			username: username,
			password: password,
		}
		fmt.Printf("%#v\n",u)
		c.JSON(http.StatusOK,gin.H{
			"message": "ok",
		})
	})
	r.Run(":9090")
}