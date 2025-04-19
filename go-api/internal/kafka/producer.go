package kafka

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"go-api/internal/model"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(broker, topic string) *Producer {
	return &Producer{
		writer: &kafka.Writer{
			Addr:         kafka.TCP(broker),
			Topic:        topic,
			Balancer:     &kafka.LeastBytes{},
			RequiredAcks: kafka.RequireOne,
		},
	}
}

func (p *Producer) SendChargeMessage(msg model.ChargeMessage) error {
	payload, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = p.writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(msg.ReferenceID),
		Value: payload,
	})
	if err != nil {
		log.Println("Kafka send error:", err)
	}
	return err
}

func (p *Producer) Close() error {
	return p.writer.Close()
}
