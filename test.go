package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"fmt"
)

func main(){
	router := gin.Default()
	router.MaxMultipartMemory = 8<<20 //8 MiB
	router.POST("/upload", func(context *gin.Context) {
		file,_ := context.FormFile("file")
		log.Println(file.Filename)
		dst := "D:\\maven\\project\\crawler\\src\\img"
		context.SaveUploadedFile(file,dst)
		log.Println(dst)
		context.String(http.StatusOK,fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	router.Run(":8080")
}
