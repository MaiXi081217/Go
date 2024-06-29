package main

import (
	"fmt"
	"html/template"
	"net/http"
	_ "github.com/gin-gonic/gin"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("parse template failed,err:%v\n", err)
		return
	}
	msg := "this index html"
	t.Execute(w, msg)

}
func home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./home.tmpl")
	if err != nil {
		fmt.Printf("parse template failed,err:%v\n", err)
		return
	}
	msg := "this home html"
	t.Execute(w, msg)
}

func index2(w http.ResponseWriter, r *http.Request){
	t,err := template.ParseFiles("./templates/base.tmpl","./templates/index2.tmpl")
	if err != nil {
		fmt.Printf("parse template failed,err:%v\n", err)
		return
	}
	msg := "this is index2"
	t.ExecuteTemplate(w,"index2.tmpl",msg)
}
func home2(w http.ResponseWriter, r *http.Request){
	t,err := template.ParseFiles("./templates/base.tmpl","./templates/home2.tmpl")
	if err != nil {
		fmt.Printf("parse template failed,err:%v\n", err)
		return
	}
	msg := "this is home2"
	t.ExecuteTemplate(w,"home2.tmpl",msg)
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home2", home2)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("HTTP start error")
		return
	}
}
