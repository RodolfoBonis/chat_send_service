package usecases

import "chat_api/infra/repositories"

type SendMessageUC struct {
	repository repositories.MessageRepository
}

func (uc *SendMessageUC) SendMessage(message string) (bool, error) {
	result, err := uc.repository.SendMessage(message)
	if err != nil {
		return false, err
	}

	return result, nil
}
