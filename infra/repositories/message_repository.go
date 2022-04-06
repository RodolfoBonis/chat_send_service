package repositories

import (
	"github.com/streadway/amqp"

	"sender_service/infra/services"
	"sender_service/models"
	"sender_service/utils"
)

type MessageRepository struct{}

func (m *MessageRepository) SendMessage(message models.MessageModel) (bool, error) {
	go func() {
		amqpService := services.AmqpService{
			UrlConnection: utils.GetEnv("AMQP_URL", "amqp://guest:guest@localhost:5672/"),
			QueueName:     "chat_messages",
		}

		channel := amqpService.OpenAmqpConnection()

		message := amqp.Publishing{
			ContentType: "application/json",
			Body:        message.ToJSON(),
		}

		err := channel.Publish(
			"chat",
			"/chat_messages",
			false,
			false,
			message,
		)

		err = channel.Close()
		if err != nil {
			return
		}
	}()

	return true, nil

}
