package producer

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

func Producer() {
	// Configure Kafka producer
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	// Create a new Kafka producer
	brokers := []string{"localhost:9092"} // Kafka broker address
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer producer.Close()

	// Message to send to Kafka
	message := &sarama.ProducerMessage{
		Topic: "example-topic",
		Value: sarama.StringEncoder("Hello, Kafka! This is a message from Go."),
	}

	// Send the message
	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		log.Fatalf("Failed to send message to Kafka: %v", err)
	}

	fmt.Printf("Message sent to partition %d with offset %d\n", partition, offset)
}
