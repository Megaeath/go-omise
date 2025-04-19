package consumer

import (
    "context"
    "fmt"
    "log"

    "github.com/segmentio/kafka-go"
)

func StartKafkaConsumer(brokerAddr, topic, groupID string, workerCount int) error {
    r := kafka.NewReader(kafka.ReaderConfig{
        Brokers:  []string{brokerAddr},
        Topic:    topic,
        GroupID:  groupID,
        MinBytes: 10e3, // 10KB
        MaxBytes: 10e6, // 10MB
    })

    defer r.Close()
    log.Println("Kafka consumer started...")

    workerPool := make(chan struct{}, workerCount)

    for {
        m, err := r.ReadMessage(context.Background())
        if err != nil {
            log.Printf("error reading message: %v", err)
            return fmt.Errorf("error reading message: %w", err) // Return error
        }

        workerPool <- struct{}{} // acquire slot

        go func(msg kafka.Message) {
            defer func() { <-workerPool }() // release slot

            ProcessChargeMessage(msg.Value)
        }(m)
    }
}