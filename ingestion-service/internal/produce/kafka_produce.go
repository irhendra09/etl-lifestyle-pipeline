package produce

import (
	"context"
	"encoding/json"
	"log"

	"github.com/irhendra09/ingestion-service/internal/model"
	"github.com/segmentio/kafka-go"
)

var writer = &kafka.Writer{
	Addr:     kafka.TCP("localhost:9092"),
	Topic:    "data-lifestyle",
	Balancer: &kafka.LeastBytes{},
}

func Produce(post *model.Post) {

	jsonData, err := json.Marshal(post)
	if err != nil {
		log.Print(err.Error())
	}

	message := kafka.Message{
		Key:   []byte("post"),
		Value: jsonData,
	}

	err = writer.WriteMessages(context.Background(), message)
	if err != nil {
		log.Print(err.Error())
	}
}
