package kafka

import (
	"context"
	"errors"
	"io"
	"log"

	"github.com/LeoUraltsev/notification-service/internal/application/service"
	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	service *service.NotificationService
	reader  *kafka.Reader
}

func New(reader *kafka.Reader, service *service.NotificationService) *KafkaConsumer {
	return &KafkaConsumer{
		service: service,
		reader:  reader,
	}
}

func (k *KafkaConsumer) Handle(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			msg, err := k.reader.FetchMessage(ctx)
			if err != nil {
				if errors.Is(err, io.EOF) {
					return
				}
				log.Println("Error: failed read message", err)
				continue
			}

			err = k.processMessage(msg)
			if err != nil {
				log.Println("Error: failed process message", err)
				continue
			}

			k.reader.CommitMessages(ctx, msg)
		}
	}
}

func (k *KafkaConsumer) processMessage(msg kafka.Message) error {
	if len(msg.Value) == 0 {
		log.Println("Warn: empty message")
	}

	err := k.service.SendNotificationCreateNewUser(string(msg.Value))
	if err != nil {
		return err
	}
	log.Println(string(msg.Value))

	return nil
}
