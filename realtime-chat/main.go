package realtime_chat

import (
	"github.com/gin-gonic/gin"
	"io"
	"fmt"
	"math/rand"
	"net/http"
)

func main(){
	router := gin.Default()
	router.SetHTMLTemplate(html)

	router.GET("/room/:roomid",roomGET)
	router.POST("/room/:roomid",roomPOST)
	router.DELETE("room/:roomid",roomDELETE)
	router.GET("/stream/:roomid",stream)

	router.Run(":8080")

}

func stream(c *gin.Context){
	roomid := c.Param("roomid")
	listener := openListener(roomid)
	defer closeListener(roomid, listener)

	clientGone := c.Writer.CloseNotify()
	c.Stream(func(w io.Writer) bool{
		select {
		case <- clientGone:
			return false
		case message := <-listener:
			c.SSEvent("message",message)
			return true
		}
	})
}

func roomGET(c *gin.Context){
	roomid := c.Param("roomid")
	userid := fmt.Sprint(rand.Int31())
	c.HTML(http.StatusOK,"chat_room",gin.H{
		"roomid":roomid,
		"userid":userid,
	})
}

func roomPOST(c *gin.Context){
	roomid := c.Param("roomid")
	userid := c.PostForm("user")
	message := c.PostForm("message")
	room(roomid).Submit(userid + ": " + message)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": message,
	})
}

func roomDELETE(c *gin.Context){
	roomid := c.Param("roomid")
	deleteBroadcast(roomid)

}
