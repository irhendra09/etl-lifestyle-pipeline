package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "data-lifestyle",
		GroupID: "etl-consumer-group",
	})

	defer reader.Close()

	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			panic(err)
		}

		fmt.Println("Key:", string(message.Key))
		fmt.Println("Value:", string(message.Value))
		fmt.Println("Partition:", message.Partition)
		fmt.Println("Offset:", message.Offset)
	}
}
