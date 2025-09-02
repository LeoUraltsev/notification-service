package domain

import "errors"

type Message struct {
	id   int64
	text string
}

func NewMessage(id int64, msg string) (*Message, error) {
	if len(msg) == 0 {
		return nil, errors.New("empty message")
	}
	return &Message{
		id:   id,
		text: msg,
	}, nil
}

func (m *Message) Text() (string, error) {
	return m.text, nil
}
