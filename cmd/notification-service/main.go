package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/LeoUraltsev/notification-service/internal/application/handlers"
	"github.com/LeoUraltsev/notification-service/internal/application/service"
	"github.com/LeoUraltsev/notification-service/internal/config"
	appkafka "github.com/LeoUraltsev/notification-service/internal/infra/kafka"
	"github.com/LeoUraltsev/notification-service/internal/infra/storage/mock"
	"github.com/LeoUraltsev/notification-service/internal/infra/telegram"
	"github.com/segmentio/kafka-go"
)

func main() {

	cfg := config.New()

	ctx, done := signal.NotifyContext(context.Background(), os.Interrupt)
	defer done()
	kafkaTopik := "test"
	kafkaConsumerGroup := "my-test"
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:29092"},
		GroupID: kafkaConsumerGroup,
		Topic:   kafkaTopik,
	})

	defer reader.Close()

	bot, err := telegram.New(cfg)
	if err != nil {
		return
	}

	repo := mock.New()
	handlerService := handlers.New(bot, repo)

	updateHandler := telegram.NewUpdateHandler(bot, handlerService)

	go func() {
		if err := updateHandler.Start(ctx); err != nil {
			return
		}
	}()

	service := service.New(bot, repo)

	k := appkafka.New(reader, service)
	go k.Handle(ctx)

	<-ctx.Done()

}
