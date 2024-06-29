package fun

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"html/template"
	"fmt"
)

func parse_tmp_err(err error){
	fmt.Printf("parse template failed,err:%v\n",err)
}

type User struct{
	Name string
	Age int
	Gender string
}

func SayGet(c *gin.Context){
	c.JSON(200,gin.H{
		"message":"Get",
	})
}

func SayPut(c *gin.Context){
	c.JSON(200,gin.H{
		"message":"Put",
	})
}

func SayDelete(c *gin.Context){
	c.JSON(200,gin.H{
		"message":"Delete",
	})
}

func SayPost(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"message":"Post",
	})
}

func SayHello(w http.ResponseWriter,r *http.Request){

	t,err := template.New(".hello.tmpl").ParseFiles("F:/GoProject/src/gin_web/templates/.hello.tmpl")
	if err != nil{
		parse_tmp_err(err)
	}

	user1 := User{
		Name: "Tom",
		Age: 18,
		Gender: "man",
	}
	
	m1 := map[string]interface{}{
		"name": "jerry",
		"age": 20,
		"gender":"man",
	}

	err = t.Execute(w,map[string]interface{}{
		"user1":user1,
		"m1":m1,
	})
	if err != nil {
		fmt.Println("reder error")
		return
	}
}

