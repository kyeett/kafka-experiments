package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/kyeett/kafka-experiments/types"
)

type exampleConsumerGroupHandler struct{}

func (exampleConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (exampleConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h exampleConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d message:%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Value)
		// sess.MarkMessage(msg, "")
		var m types.MyMessage
		if err := json.Unmarshal(msg.Value, &m); err != nil {
			fmt.Println("Something went wrong", err)
			continue
		}

		fmt.Printf("%s\n", m)
	}
	return nil
}

func main() {
	// Init config, specify appropriate version
	config := sarama.NewConfig()
	config.Version = sarama.V1_0_0_0
	config.Consumer.Return.Errors = true

	// Start with a client
	client, err := sarama.NewClient([]string{"localhost:9092"}, config)
	if err != nil {
		panic(err)
	}
	defer func() { _ = client.Close() }()

	// Start a new consumer group
	group, err := sarama.NewConsumerGroupFromClient("my-group", client)
	if err != nil {
		panic(err)
	}
	defer func() { _ = group.Close() }()

	// Track errors
	go func() {
		for err := range group.Errors() {
			fmt.Println("ERROR", err)
		}
	}()

	fmt.Println("STARTING")
	// Iterate over consumer sessions.
	ctx := context.Background()
	for {
		topics := []string{"test"}
		handler := exampleConsumerGroupHandler{}

		err := group.Consume(ctx, topics, handler)
		if err != nil {
			panic(err)
		}
	}
}
