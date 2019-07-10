package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	r := gin.Default()
	r.GET("/JSONP?callback=x", func(context *gin.Context) {
		data := map[string]interface{}{
			"foo":"bar",
		}
		context.JSONP(http.StatusOK,data)
	})
	r.Run(":8080")
}
