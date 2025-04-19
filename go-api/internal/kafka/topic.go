// internal/kafka/topic.go
package kafka

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func EnsureTopic(brokerAddress, topic string, partitions int, replicationFactor int) error {
	conn, err := kafka.Dial("tcp", brokerAddress)
	if err != nil {
		return fmt.Errorf("failed to connect to broker: %w", err)
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		return fmt.Errorf("failed to get controller: %w", err)
	}

	ctrlConn, err := kafka.Dial("tcp", fmt.Sprintf("%s:%d", controller.Host, controller.Port))
	if err != nil {
		return fmt.Errorf("failed to connect to controller: %w", err)
	}
	defer ctrlConn.Close()

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Pass the topic configuration as individual arguments
	if err := ctrlConn.CreateTopics(kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     partitions,
		ReplicationFactor: replicationFactor,
	}); err != nil {
		return fmt.Errorf("failed to create topic: %w", err)
	}

	return nil
}
