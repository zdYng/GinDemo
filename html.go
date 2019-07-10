package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")
	router.GET("/posts/index", func(context *gin.Context) {
		context.HTML(http.StatusOK,"posts/index.tmpl",gin.H{
			"title":"Posts",
		})
	})
	router.GET("/users/index", func(context *gin.Context) {
		context.HTML(http.StatusOK,"users/index.tmpl",gin.H{
			"title":"Users",
		})
	})
	router.Run(":8080")
}
