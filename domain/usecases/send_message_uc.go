package usecases

import (
	"sender_service/infra/repositories"
	"sender_service/models"
)

type SendMessageUC struct {
	repository repositories.MessageRepository
}

func (uc *SendMessageUC) SendMessage(message models.MessageModel) (bool, error) {

	result, err := uc.repository.SendMessage(message)
	if err != nil {
		return false, err
	}

	return result, nil
}
