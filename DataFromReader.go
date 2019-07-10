package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	r := gin.Default()
	r.GET("/someDataFromReader", func(context *gin.Context) {
		response,err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || response.StatusCode != http.StatusOK{
			context.Status(http.StatusServiceUnavailable)
			return
		}
		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition":`attachment; filename="gopher.png"`,
		}
		context.DataFromReader(http.StatusOK,contentLength,contentType,reader,extraHeaders)
	})
	r.Run(":8085")
}
