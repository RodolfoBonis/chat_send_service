package routes

import (
	"github.com/gin-gonic/gin"

	"sender_service/domain/usecases"
	"sender_service/models"
)

func messageRoutes(r *gin.Engine) {
	user := r.Group("/message")
	{
		user.POST("/send", sendMessage)
	}
}

func sendMessage(c *gin.Context) {
	usecase := usecases.SendMessageUC{}
	var message models.MessageModel

	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err := usecase.SendMessage(message)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"message": "message sent",
	})
}
