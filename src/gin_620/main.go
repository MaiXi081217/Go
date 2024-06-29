package main

import (
	"fmt"
 	"html/template"
	"net/http"
	_ "github.com/gin-gonic/gin"
)

func SayHello(w http.ResponseWriter,r *http.Request){

	shuChu := func (arg string) (string,error)  {
		return arg + "fuckoff",nil
	}


	t := template.New("test.tmpl") 
	t.Funcs(template.FuncMap{		//告诉模板引擎添加一个自定义函数
		"shuChu":shuChu,			//前面为模板文件调用时用的名字，后为实际定义函数名
	})


	_ , err := t.ParseFiles("./test.tmpl")
	if err != nil{
		fmt.Printf("parse template error,err:%v\n",err)
	}
	name := "詹文艺"
	err = t.Execute(w,name)
	if err != nil{
		fmt.Printf("reder error,err:%V\n",err)
	}
}
func demo01(w http.ResponseWriter,r *http.Request){
	t,err := template.ParseFiles("./t.tmpl","./ul.tmpl")
	if err != nil{
		fmt.Printf("parse failed err:%v\n",err)
	}
	name := "mm"

	if err != nil{
		fmt.Printf("parse template failed,err:%v\n",err)
	}
	err = t.Execute(w,name)
	if err != nil {
		fmt.Printf("reder failed,err:%v\n",err)
	}

}

func main() {
http.HandleFunc("/",SayHello)
http.HandleFunc("/demo",demo01)
err := http.ListenAndServe(":8080",nil)
if err != nil{
	fmt.Println("HTTP start error")
	return
	}
}