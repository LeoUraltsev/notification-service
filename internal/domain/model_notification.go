package domain

import "errors"

type ChatID struct {
	value int64
}

func NewChatID(id int64) (ChatID, error) {
	return ChatID{
		value: id,
	}, nil
}

func (c ChatID) Value() int64 {
	return c.value
}

type Message struct {
	value string
}

func NewMessage(msg string) (Message, error) {
	return Message{
		value: msg,
	}, nil
}

func (m Message) Value() string {
	return m.value
}

type Notification struct {
	id      ChatID
	message Message
}

func NewNotification(id ChatID, msg Message) (*Notification, error) {
	if len(msg.value) == 0 {
		return nil, errors.New("empty message")
	}
	if len(msg.value) > 4096 {
		return nil, errors.New("long message text")
	}
	return &Notification{
		id:      id,
		message: msg,
	}, nil
}

func (m *Notification) Message() Message {
	return m.message
}

func (m *Notification) ID() ChatID {
	return m.id
}
