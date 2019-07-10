package main
//绑定uri
import "github.com/gin-gonic/gin"

type Person struct {
	ID string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:required`
}

type myForm struct {
	Colors []string `form:"colors[]"`
}

type LoginForm struct {
	User string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main(){
	router := gin.Default()
	//router.LoadHTMLGlob("templates/*")
	router.GET("/:name/:id", func(context *gin.Context) {
		var person Person
		if err := context.ShouldBindUri(&person);err!=nil{
			context.JSON(400,gin.H{"msg":err})
			return
		}
		context.JSON(200,gin.H{"name":person.Name,"uuid":person.ID})
	})

	router.POST("/",formHandler)
	router.GET("/",indexHandler)
	router.POST("/login", func(context *gin.Context) {
		var form LoginForm
		if context.ShouldBind(&form) == nil{
			if form.User == "admin"&& form.Password=="admin123"{
				context.JSON(200,gin.H{"status":"you are logged in"})
			}else{
				context.JSON(401,gin.H{"status": "unauthorized"})
			}
		}
	})

	router.Run(":8085")
}

func indexHandler(c *gin.Context){
	c.HTML(200,"form.html", nil)
}

func formHandler(context *gin.Context){
	var fakeForm myForm
	context.Bind(&fakeForm)
	context.JSON(200, gin.H{"color": fakeForm.Colors})
}
