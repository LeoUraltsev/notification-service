package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/LeoUraltsev/notification-service/internal/application"
	"github.com/LeoUraltsev/notification-service/internal/client/telegram"
	"github.com/LeoUraltsev/notification-service/internal/config"
	appkafka "github.com/LeoUraltsev/notification-service/internal/infra/kafka"
	"github.com/LeoUraltsev/notification-service/internal/infra/storage/mock"
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

	go bot.Start()

	repo := mock.New()
	service := application.New(bot, repo)

	k := appkafka.New(reader, service)
	go k.Handle(ctx)

	<-ctx.Done()

}
