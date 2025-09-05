package handlers

import (
	"context"
	"fmt"

	"github.com/LeoUraltsev/notification-service/internal/domain"
)

type Sender interface {
	SendRequestPhoneNumber(notification *domain.Notification) error
	SendNotification(notification *domain.Notification) error
}

type TelegramCommands struct {
	sender   Sender
	userRepo domain.UserRepository
}

func New(sender Sender, userRepo domain.UserRepository) *TelegramCommands {
	return &TelegramCommands{
		sender:   sender,
		userRepo: userRepo,
	}
}

func (t *TelegramCommands) HandleStartCommand(ctx context.Context, chatID domain.ChatID) error {
	msg := "Привет! Этот бот создан для получения уведомлений от сервиса авторизации.\n\nОтправь номер телефона для продолжения."
	message, err := domain.NewMessage(msg)
	if err != nil {
		return err
	}
	notification, err := domain.NewNotification(chatID, message)
	if err != nil {
		return err
	}
	err = t.sender.SendRequestPhoneNumber(notification)
	if err != nil {
		return err
	}

	return nil
}

func (t *TelegramCommands) HandleGetPhoneCommand(ctx context.Context, chatID domain.ChatID, phone string) error {
	msg := fmt.Sprintf("Получили ваш номер телефона: %s", phone)
	message, err := domain.NewMessage(msg)
	if err != nil {
		return err
	}
	notification, err := domain.NewNotification(chatID, message)
	if err != nil {
		return err
	}
	t.sender.SendNotification(notification)

	return nil
}
