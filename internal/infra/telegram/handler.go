package telegram

import (
	"context"
	"fmt"

	"github.com/LeoUraltsev/notification-service/internal/application/handlers"
	"github.com/LeoUraltsev/notification-service/internal/domain"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	CommandStart = "start"
)

type UpdatesHandler struct {
	client  *Client
	handler *handlers.TelegramCommands
}

func NewUpdateHandler(client *Client, handler *handlers.TelegramCommands) *UpdatesHandler {
	return &UpdatesHandler{
		client:  client,
		handler: handler,
	}
}

func (h *UpdatesHandler) Start(ctx context.Context) error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := h.client.bot.GetUpdatesChan(u)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("контекст остановлен")
			return ctx.Err()
		case update := <-updates:
			if err := h.handleUpdate(ctx, update); err != nil {
				return err
			}
		}
	}

}

func (h *UpdatesHandler) handleUpdate(ctx context.Context, update tgbotapi.Update) error {
	if update.Message == nil {
		return nil
	}

	id := update.Message.Chat.ID
	chatID, err := domain.NewChatID(id)
	if err != nil {
		return err
	}
	switch {
	case update.Message.Command() == CommandStart:
		return h.handler.HandleStartCommand(ctx, chatID)
	case update.Message.Contact != nil:
		phone := update.Message.Contact.PhoneNumber
		return h.handler.HandleGetPhoneCommand(ctx, chatID, phone)
	}

	return nil
}
