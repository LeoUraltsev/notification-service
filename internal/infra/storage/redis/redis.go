package redis

import (
	"context"

	"github.com/LeoUraltsev/notification-service/internal/domain"
)

type Storage struct {
	redis Client
}

func NewStorage(redis Client) *Storage {
	return &Storage{
		redis: redis,
	}
}

func (s *Storage) AdminUsers(ctx context.Context) ([]*domain.User, error) {

	return nil, nil
}
func (s *Storage) UserByID(ctx context.Context, id domain.UserID) (*domain.User, error) {

	return nil, nil
}
