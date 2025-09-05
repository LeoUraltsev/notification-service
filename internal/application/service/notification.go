package service

import (
	"errors"
	"fmt"

	"github.com/LeoUraltsev/notification-service/internal/domain"
)

type Sender interface {
	SendNotification(notification *domain.Notification) error
}

type NotificationService struct {
	Sender Sender
	Repo   domain.UserRepository
}

func New(sender Sender, repo domain.UserRepository) *NotificationService {
	return &NotificationService{
		Sender: sender,
		Repo:   repo,
	}
}

func (s *NotificationService) SendNotificationCreateNewUser(msg string) error {

	if len(msg) == 0 {
		return errors.New("message is empty")
	}

	message, err := domain.NewMessage(msg)
	if err != nil {
		fmt.Println("ошибка валидации сообщения", msg)
	}

	users, err := s.Repo.AdminUsers()
	if err != nil {
		return err
	}

	for _, cID := range users {
		chatID, err := domain.NewChatID(cID)
		if err != nil {
			fmt.Println("неверный id", chatID)
			continue
		}

		notification, err := domain.NewNotification(chatID, message)
		if err != nil {
			return err
		}
		err = s.Sender.SendNotification(notification)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	return nil
}
