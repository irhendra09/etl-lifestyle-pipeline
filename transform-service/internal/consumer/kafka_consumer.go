package consumer

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

var reader = kafka.NewReader(kafka.ReaderConfig{
	Brokers: []string{"localhost:9092"},
	Topic:   "data-lifestyle",
	GroupID: "etl-consumer-group",
})

func StartConsumer(handler func(kafka.Message) error) {
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println(err)
			continue
		}

		err = handler(msg)
		if err != nil {
			log.Println(err)
		}
	}
}
