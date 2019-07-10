package main

import (
	"github.com/gin-gonic/gin"
	//"github.com/gin-gonic/autotls"
	"time"
	"log"
	"net/http"
)

func main(){
	r := gin.Default()

	r.GET("long_async", func(c *gin.Context) {
		cCp := c.Copy()
		go func() {
			time.Sleep(5 *time.Second)
			log.Println("Done ! in path" + cCp.Request.URL.Path)
		}()
	})
	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		log.Println("Done ! in path "+ c.Request.URL.Path)
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(200,"pong")
	})
	log.Fatal(autotls)
	//r.Run(":8085")
	http.ListenAndServe(":8085",r) // 自定义HTTP配置
	s := &http.Server{
		Addr: ":8085",
		Handler: r,
		ReadTimeout: 3*time.Second,
		WriteTimeout: 3*time.Second,
		MaxHeaderBytes: 1<<20,
	}
	s.ListenAndServe()
}
