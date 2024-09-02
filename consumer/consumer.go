package consumer

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

func Consumer() {
	// Configure Kafka consumer
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Create a new Kafka consumer
	brokers := []string{"localhost:9092"} // Kafka broker address
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}
	defer consumer.Close()

	// Consume messages from the topic
	topic := "example-topic"
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to consume from Kafka topic: %v", err)
	}
	defer partitionConsumer.Close()

	// Listen for messages
	fmt.Println("Listening for messages...")
	for message := range partitionConsumer.Messages() {
		fmt.Printf("Received message: %s\n", string(message.Value))
	}
} 

}
