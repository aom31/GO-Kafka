package producer

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func producer() {
	servers := []string{"localhost:9092"}

	producer, err := sarama.NewSyncProducer(servers, nil)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	msg := sarama.ProducerMessage{
		Topic: "aomhello",
		Value: sarama.StringEncoder("hello workd"),
	}

	partitionRes, offsetRes, err := producer.SendMessage(&msg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("partition=%v , offset=%v", partitionRes, offsetRes)
}
