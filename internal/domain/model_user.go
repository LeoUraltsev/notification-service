package domain

import "github.com/google/uuid"

type UserID struct {
	value uuid.UUID
}

func NewUserID(id uuid.UUID) (UserID, error) {
	return UserID{value: id}, nil
}

type PhoneNumber struct {
	value string
}

func NewPhoneNumber(phone string) (PhoneNumber, error) {
	return PhoneNumber{
		value: phone,
	}, nil
}

type User struct {
	userID  UserID
	chatID  ChatID
	phone   PhoneNumber
	isAdmin bool
}

func NewUser(userID UserID, chatID ChatID, phone PhoneNumber, isAdmin bool) (*User, error) {

	return &User{
		userID:  userID,
		chatID:  chatID,
		isAdmin: isAdmin,
	}, nil
}
