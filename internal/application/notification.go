package application

import (
	"fmt"

	"github.com/LeoUraltsev/notification-service/internal/domain"
)

type Sender interface {
	SendNotification(chatID int64, message string) error
}

type NotificationService struct {
	Sender Sender
	Repo   domain.Repository
}

func New(sender Sender, repo domain.Repository) *NotificationService {
	return &NotificationService{
		Sender: sender,
		Repo:   repo,
	}
}

func (s *NotificationService) SendNotificationCreateNewUser(msg string) error {

	users, err := s.Repo.AdminUsers()
	if err != nil {
		return err
	}

	for _, user := range users {
		err := s.Sender.SendNotification(user, msg)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	return nil
}
