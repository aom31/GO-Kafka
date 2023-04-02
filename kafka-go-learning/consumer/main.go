package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func main() {
	server := []string{"localhost:9092"}
	consumer, err := sarama.NewConsumer(server, nil)
	if err != nil {
		panic(err)
	}

	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("aomhello", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	defer partitionConsumer.Close()
	//loop for get message from topic we subscript
	for {
		select {
		case err := <-partitionConsumer.Errors():
			fmt.Println(err)
		case msg := <-partitionConsumer.Messages():
			fmt.Println(string(msg.Value))
		}
	}
}
