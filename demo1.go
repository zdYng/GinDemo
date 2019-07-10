package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
)

func main(){
	router := gin.Default() // 包含Logger、Recovery中间件
	// 禁用控制台颜色
	//gin.DisableConsoleColor()
	//
	// 创建记录日志的文件
	//f,_ := os.Create("testdata/gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	// 如需要将日志同时写入文件和控制台，请使用如下代码
	//gin.DefaultWriter = io.MultiWriter(f,os.Stdout)
	//router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK,"Hello %s",name)
	})

	router.GET("/user/:name/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")
		message := name + " is " + action
		context.String(http.StatusOK,message)
	})

	// 获取 Get 参数
	router.GET("/welcome", func(context *gin.Context) {
		firstname := context.DefaultQuery("firstname", "Guest")
		lastname := context.Query("lastname")
		lastname2 := context.Request.URL.Query().Get("lastname")
		context.String(http.StatusOK,"Hello %s | %s | %s",firstname,lastname,lastname2)
	})

	// 获取 POST 参数
	router.POST("/form_post", func(context *gin.Context) {
		message := context.PostForm("message")
		nick := context.DefaultQuery("nick","anonymous")
		context.JSON(200, gin.H{
			"status": "posted",
			"message": message,
			"nick": nick,
		})
	})

	// GET + POST 混合
	router.POST("/post", func(context *gin.Context) {
		id := context.Query("id")
		page := context.DefaultQuery("page","0")
		name := context.PostForm("name")
		message := context.PostForm("message")

		//fmt.Printf("id: %s; page:%s;name:%s;message:%s",id,page,name,message)
		context.JSON(200,gin.H{
			"status":200,
			"message":message,
			"name":name,
			"page":page,
			"id":id,
		})
	})

	// 上传文件(单个)
	//router := gin.Default()
	router.MaxMultipartMemory = 8<< 20
	router.POST("/upload", func(context *gin.Context) {
		file,_ := context.FormFile("file")
		log.Println(file.Filename)
		context.SaveUploadedFile(file,"testdata/kk.dll")
		context.JSON(200,gin.H{
			"status": 200,
			"message": "upload ok",
			"filename": file.Filename,
		})
	})

	// 多个文件上传
	router.POST("/uploads", func(context *gin.Context) {
		form,_ := context.MultipartForm()
		files := form.File["upload[]"]
		log.Println(context.Request.Form,files)
		for _, file := range files{
			dst := "testdata/" + file.Filename
			log.Println(file.Filename)
			context.SaveUploadedFile(file,dst)
		}
		context.JSON(200,gin.H{
			"status":200,
			"file": len(files),
		})
	})

	// 路由分组
	//v1 := router.Group("/v1", gin.Recovery())
	//v1.Use(gin.Logger())
	//{
	//	v1.POST("/login", gin.Recovery(),loginEndpoint)
	//	v1.POST("/submit", submitEndpoint)
	//	v1.POST("/read", readEndpoint)
	//}
	//v2 := router.Group("/v2")
	//{
	//	v2.POST("/login",loginEndpoint)
	//	v2.POST("/submit",submitEndpoint)
	//	v2.POST("/read", readEndpoint)
	//}

	router.Run("192.168.1.104:8080")
}
