package main
import(
	"github.com/gin-gonic/gin"
	"net/http"
)

func hello(c *gin.Context){
	c.HTML(http.StatusOK,"index.tmpl",gin.H{
		""
	})
}

func main(){
	r := gin.Default()

	r.LoadHTMLFiles("./templates/index.tmpl")  // 模板解析

	r.GET("/index",func(c *gin.Context){
		c.HTML(http.StatusOK,"index.tmpl",gin.H{  // 模板渲染
			"title":"hello,Mc",
		})
	})

	r.Run(":9090")


}