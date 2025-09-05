package telegram

import (
	"github.com/LeoUraltsev/notification-service/internal/config"
	"github.com/LeoUraltsev/notification-service/internal/domain"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Client struct {
	bot *tgbotapi.BotAPI
}

func New(cfg *config.Config) (*Client, error) {

	bot, err := tgbotapi.NewBotAPI(cfg.Telegram.Token)
	if err != nil {
		return nil, err
	}

	bot.Debug = cfg.Telegram.Debug

	return &Client{bot: bot}, nil
}

func (c *Client) SendNotification(notification *domain.Notification) error {
	msg := tgbotapi.NewMessage(notification.ID().Value(), notification.Message().Value())
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	c.bot.Send(msg)
	return nil
}

func (c *Client) SendRequestPhoneNumber(notification *domain.Notification) error {

	msg := tgbotapi.NewMessage(notification.ID().Value(), notification.Message().Value())

	keyboard := tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButtonContact("Отправить номер телефона"),
	))

	msg.ReplyMarkup = keyboard

	c.bot.Send(msg)

	return nil
}
