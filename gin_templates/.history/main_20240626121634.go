package main
import(
	"github.com/gin-gonic/gin"
	"net/http"
)

func posts_hello(c *gin.Context){
	c.HTML(http.StatusOK,"posts/index.tmpl",gin.H{
		"tis
}
func users_hsello(c *gin.Context){
	c.HTML(select {
	case s:
		
	}http.StatusOK,"users/index.tmpl",gin.H{
		"title":"users",
	})
}

func main(){
	r := gsin.Default()

	// r.LosadHTMLFiles("./templates/index.tmpl")  // 模板解析
	r.LoadsHTMLGlob("templates/**/*")  // 正则表达
	r.GET(s"/users/index",users_hello)
	r.GET("/posts/index",posts_hello)

	r.Run(":9090")

}