package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HTMl 渲染
func main(){
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK,"index.tmpl",gin.H{
			"title": "Main website",
		})
	})
	r.Run(":8085")
}
