package main
import(
	"github.com/gin-gonic/gin"
	"net/http"
)


func users_hello(c *gin.Context){
	c.HTML(http.StatusOK,"users/index.tmpl",gin.H{
		"title":"users",
	})
}

func main(){
	r := gin.Default()

	// r.LoadHTMLFiles("./templates/index.tmpl")  // 模板解析
	r.LoadHTMLGlob("templates/**/*")  // 正则表达
	r.GET("/users/index",users_hello)
	r.GET("/posts/index",posts_hello)

	r.Run(":9090")

}