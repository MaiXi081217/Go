package main
import(
	"github.com/gin-gonic/gin"
	"net/http"
)

func post_hello(c *gin.Context){
	c.HTML(http.StatusOK,"posts/index.tmpl",gin.H{
		"title":"dmh",
	})
}
func users_hello(c *gin.Context){
	c.HTML(http.StatusOK,"posts/index.tmpl",gin.H{
		"title":"dmh",
	})
}

func main(){
	r := gin.Default()

	r.LoadHTMLFiles("./templates/index.tmpl")  // 模板解析
	r.LoadHTMLGlob("template/**/*")  // 正则表达
	r.GET("/users/index",hello)
	r.GET("/posts/index",hello)

	r.Run(":9090")

}