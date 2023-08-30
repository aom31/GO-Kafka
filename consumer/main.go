package main

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func main() {
	serversAddr := []string{"localhost:9092"}
	//1. create consumer โดยการส่่ง address server เข้าไป ซึ่งสามารถส่งได้หลายตัว ถ้าเกิดว่ามี kafka หลาบclusters
	consumer, err := sarama.NewConsumer(serversAddr, nil)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	//2. สร้างการ consume ไปที่ topic และ partition ที่ต้องการ
	const topicKafka = "learnkafka-aomtest"
	partitionConsumer, err := consumer.ConsumePartition(topicKafka, 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	defer partitionConsumer.Close()

	// 3. รอรับ message โดยเขาจะส่งเป็น channel กลับมา
	log.Println("consumer start success")
	for {
		select {
		case err := <-partitionConsumer.Errors():
			fmt.Println(err)
		case msg := <-partitionConsumer.Messages():
			fmt.Println(string(msg.Value))
		}
	}
}
