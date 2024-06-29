package main
import(
	"github.com/gin-gonic/gin"
	"net/http"
)

func hello(c *gin.Context){
	c.HTML(http.StatusOK,"index.tmpl",gin.H{
		"title":"dmh",
	})
}

func main(){
	r := gin.Default()

	r.LoadHTMLFiles("./templates/index.tmpl")  // 模板解析
	r.LoadHTMLGlob("template/**/*")  // 正则表达
	r.GET("/index",hello)
	r.GET()

	r.Run(":9090")

}