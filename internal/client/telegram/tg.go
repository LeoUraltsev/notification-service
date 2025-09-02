package telegram

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Telegram struct {
	Bot *tgbotapi.BotAPI
}

func New(token string) (*Telegram, error) {

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	bot.Debug = true

	return &Telegram{Bot: bot}, nil
}

func (t *Telegram) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30
	updates := t.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message.Text == "/start" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, `Привет! Этот бот создан для получения уведомлений от сервиса авторизации. 

Отправь номер телефона для продолжения.`)
			keyboard := tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButtonContact("Отправить номер телефона"),
			))

			msg.ReplyMarkup = keyboard

			t.Send(msg)

		}
		if update.Message.Contact != nil {
			contact := update.Message.Contact
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Спасибо за номер телефона! Ваш номер: %s", contact.PhoneNumber))
			t.Send(msg)
		}
	}

	return nil
}

func (t *Telegram) Send(msg tgbotapi.MessageConfig) error {
	t.Bot.Send(msg)
	return nil
}

func (t *Telegram) SendNotification(chatID int64, message string) error {
	msg := tgbotapi.NewMessage(chatID, message)
	t.Bot.Send(msg)
	return nil
}
