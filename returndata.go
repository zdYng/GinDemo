package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

// XML、JSON、YAML、ProtoBuf（输出格式）

func main(){
	r := gin.Default()
	r.GET("/someJSON", func(context *gin.Context) {
		context.JSON(http.StatusOK,gin.H{"message":"hey","status":http.StatusOK})
	})

	r.GET("/moreJSON", func(context *gin.Context) {
		var msg struct{
			Name string `json:"user"`
			Message string
			Number int
		}
		msg.Name = "admin"
		msg.Message = "hey"
		msg.Number = 123

		context.JSON(http.StatusOK,msg)
	})

	r.GET("/someXML", func(context *gin.Context) {
		context.XML(http.StatusOK,gin.H{"message":"hey","status":http.StatusOK})
	})
	r.GET("/someYAML", func(context *gin.Context) {
		context.YAML(http.StatusOK,gin.H{"message":"hey","status":http.StatusOK})
	})
	r.GET("/someProtoBuf", func(context *gin.Context) {
		reps := []int64{int64(1),int64(2)}
		label := "test"
		data := &protoexample.Test{
			Label:&label,
			Reps:reps,
		}
		context.ProtoBuf(http.StatusOK,data)
	})
	r.Run(":8085")

}
