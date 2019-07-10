package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	r := gin.Default()
	r.Static("/assets","assets")
	r.StaticFS("/more_static",http.Dir("my_file_system"))
	r.StaticFile("/form.html","templates/form.html")
	r.Run(":8085")
}
