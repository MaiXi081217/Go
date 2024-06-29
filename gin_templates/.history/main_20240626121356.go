package main
import(
	"github.com/gin-gonic/gin"
	"net/http"
)




func main(){
	r := gin.Default()

	// r.LoadHTMLFiles("./templates/index.tmpl")  // 模板解析
	r.LoadHTMLGlob("templates/**/*")  // 正则表达
	r.GET("/users/index",users_hello)
	r.GET("/posts/index",func(c *gin.Context){
		c.HTML(http.StatusOK,"posts/index.tmpl",gin.H{
			"title":"posts",
		})
	})

	r.Run(":9090")

}