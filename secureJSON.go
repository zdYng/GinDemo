package main
//使用SecureJSON可以防止json劫持，如果返回的数据是数组，则会默认在返回值前加上"while(1)"
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	r := gin.Default()
	r.GET("/someJSON", func(context *gin.Context) {
		names := []string{"admin","admin2","admin3"}
		context.SecureJSON(http.StatusOK,gin.H{"data":names})

	})
	r.GET("/purejson", func(context *gin.Context) {
		context.PureJSON(200,gin.H{"html":"<b>Hello,World</b>"})
	})
	r.Run(":8085")
}
