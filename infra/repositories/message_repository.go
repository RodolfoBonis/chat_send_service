package repositories

import (
	"chat_api/infra/services"
	"chat_api/utils"
	"github.com/streadway/amqp"
)

type MessageRepository struct{}

func (m *MessageRepository) SendMessage(message string) (bool, error) {
	go func() {
		amqpService := services.AmqpService{
			UrlConnection: utils.GetEnv("AMQP_URL", "amqp://guest:guest@localhost:5672/"),
			QueueName:     "chat_messages",
		}

		channel := amqpService.OpenAmqpConnection()

		message := amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
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
