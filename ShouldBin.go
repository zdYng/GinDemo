package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 绑定为json
type Login struct {
	User string `form:"user" json:"user" xml:"user" binding:"-"`
	Password string `form:"password" json:"password" xml:"password" binding:"-"`
}

func main(){
	router := gin.Default()
	router.POST("/loginJSON", func(context *gin.Context) {
		var json Login
		if err := context.ShouldBindJSON(&json);err!=nil{
			context.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
		if json.User != "admin" || json.Password!="admin123"{
			context.JSON(http.StatusUnauthorized,gin.H{"status": "unauthorized"})
			return
		}

		context.JSON(http.StatusOK,gin.H{"status":"you are logged in"})
	})

	router.POST("/loginXML", func(context *gin.Context) {
		var xml Login
		if err := context.ShouldBindXML(&xml); err!=nil{
			context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}

		if xml.User != "admin" || xml.Password != "123"{
			context.JSON(http.StatusUnauthorized,gin.H{"status":"unauthorized"})
			return
		}
		context.JSON(http.StatusOK,gin.H{"status":"you are logged in"})
	})

	router.POST("/loginForm", func(context *gin.Context) {
		var form Login
		if err := context.ShouldBind(&form); err!=nil{
			context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		if form.User != "admin" || form.Password!="admin123"{
			context.JSON(http.StatusUnauthorized,gin.H{"status":"unauthorized"})
			return
		}
		context.JSON(http.StatusOK,gin.H{"status":"you are logged in"})
	})

	router.Run("0.0.0.0:9090")
}

