package main

import (
	"log"

	"github.com/IBM/sarama"
)

func main() {

	serversAddr := []string{"localhost:9092"}
	//1. create producer with server
	syncProducer, err := sarama.NewSyncProducer(serversAddr, nil)
	if err != nil {
		panic(err)
	}
	defer syncProducer.Close()

	//2.0 เตรียม producer message ก่อนส่่ง
	const topicKafka = "learnkafka-aomtest"
	msgProduce := sarama.ProducerMessage{
		Topic: topicKafka,
		Value: sarama.StringEncoder("test msg"),
	}

	//2. producer send message to consumer
	partitionSend, offsetSend, err := syncProducer.SendMessage(&msgProduce)
	if err != nil {
		panic(err)
	}
	log.Printf("partition send msg:%v, offset:%v", partitionSend, offsetSend)
}
