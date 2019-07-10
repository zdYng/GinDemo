package main

import (
	"github.com/gin-gonic/gin"
	"testing"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"net/http"
)

func setupRouter() *gin.Engine{
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200,"pong")
	})
	return r
}

func main(){
	r := setupRouter()
	r.Run(":8086")
}

func TestPingRouter(t *testing.T){
	router := setupRouter()
	w := httptest.NewRecorder()
	req,_ := http.NewRequest("GET","/ping",nil)
	router.ServeHTTP(w,req)
	assert.Equal(t,200,w.Code)
	assert.Equal(t,"pong",w.Body.String())
}