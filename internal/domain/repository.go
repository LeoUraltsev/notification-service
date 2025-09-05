package domain

import (
	"context"
)

type UserRepository interface {
	AdminUsers() ([]int64, error)
}

type UserRepositoryNew interface {
	AdminUsers() ([]*User, error)
	UserByID(id UserID) (*User, error)
	UserByChatID(id ChatID) (*User, error)
	UserByPhoneNumber(phone PhoneNumber) (*User, error)
}

type NotificationSender interface {
	Send(ctx context.Context, notification *Notification)
}
