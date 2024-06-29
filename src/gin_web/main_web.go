package main

import (
	_"fmt"
	"gin_web/fun"
	_"html/template"
	"net/http"
	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()

	r.GET("/",fun.SayGet)
	r.POST("/",fun.SayPost)
	r.DELETE("/",fun.SayDelete)
	r.PUT("/",fun.SayPut)

	http.HandleFunc("/template",fun.SayHello)

	r.Run() 
}