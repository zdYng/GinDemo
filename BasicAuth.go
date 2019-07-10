package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// BasicAuth Middleware

var secrets = gin.H{
	"foo": gin.H{"email":"foo@bar.com","phone":"123321"},
	"gaga": gin.H{"email":"gaga@123.com", "phone":"999"},
	"llala": gin.H{"email":"llala@133.com","phone":"11222"},
}

func main(){
	r := gin.Default()
	authorized := r.Group("/admin",gin.BasicAuth(gin.Accounts{
		"foo":"bar",
		"gaga": "123",
		"llala": "haha",
	}))

	authorized.GET("/secrets", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret,ok := secrets[user];ok{
			c.JSON(http.StatusOK,gin.H{"user":user,"secret":secret})
		}else {
			c.JSON(http.StatusOK,gin.H{"user":user,"secret":"NO SECRET : ("})
		}
	})
	r.Run(":8085")
}
