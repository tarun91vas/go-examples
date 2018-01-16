package main

import (
	"fmt"
	"strconv"

	kafka "github.com/Shopify/sarama"
)

func main() {
	//Setup config
	config := kafka.NewConfig()
	config.Producer.Flush.Messages = 10
	config.Producer.Return.Successes = true

	//Get a producer object
	producer, err := kafka.NewAsyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection Successful !!")

	//Send producer messages to Input channel
	for i := 0; i < 10; i++ {
		producer.Input() <- &kafka.ProducerMessage{
			Topic: "test-topic",
			Key:   nil,
			Value: kafka.StringEncoder("Event: " + strconv.Itoa(i)),
		}
	}

	//Catch errors and success acknowledgement
	for i := 0; i < 10; i++ {
		select {
		case msg := <-producer.Errors():
			fmt.Println(msg.Err)
		case success := <-producer.Successes():
			fmt.Println(success.Topic, success.Value)
		}
	}

	//Close producer to avoid leaks
	producer.Close()
}
