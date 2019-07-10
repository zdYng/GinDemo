package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 重定向

func Redirect(c * gin.Context){
	c.Redirect(http.StatusMovedPermanently,"http://www.google.com/") //外部重定向
}



func main(){
	r := gin.Default()

	r.GET("/redirect", Redirect)

	// 路由重定向  内部
	r.GET("/test", func(context *gin.Context) {
		context.Request.URL.Path = "/test2"
		r.HandleContext(context)
	})
	r.GET("/test2", func(context *gin.Context) {
		context.JSON(200,gin.H{"hello":"world"})
	})
	r.Run(":8085")
}
