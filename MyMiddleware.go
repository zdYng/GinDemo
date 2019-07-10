package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"log"
)

// 自定义中间件
func Logger() gin.HandlerFunc{
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("example","123456")

		c.Next()
		latency := time.Since(t)

		log.Print(latency)
		status := c.Writer.Status()
		log.Println(status)
		return
	}
}

func main(){
	r := gin.New()
	r.Use(Logger())
	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		log.Println(example)
		return
	})
	r.Run(":8085")
}
