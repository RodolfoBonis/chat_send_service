package routes

import (
	"chat_api/domain/usecases"
	"github.com/gin-gonic/gin"
)

func messageRoutes(r *gin.Engine) {
	user := r.Group("/message")
	{
		user.POST("/send", sendMessage)
	}
}

func sendMessage(c *gin.Context) {
	usecase := usecases.SendMessageUC{}

	_, err := usecase.SendMessage(c.PostForm("message"))
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"message": "message sent",
	})
}
