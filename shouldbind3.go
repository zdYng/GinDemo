package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type Person struct {
	Name string `form:"name"`
	Address string `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main(){
	router := gin.Default()
	router.Any("/testing",startPage)
	router.Run(":8085")
}

func startPage(context *gin.Context){
	var person Person
	if context.ShouldBind(&person)==nil{
		log.Println("====Onlu Bind By Query String ====")
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	}
	context.String(200,"Success")
}
