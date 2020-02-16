package main

import (
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/Shopify/sarama"
)

var (
	wg                          sync.WaitGroup
	enqueued, successes, errors int
)

func main() {

	// Trap SIGINT to trigger a graceful shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		panic(err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for range producer.Successes() {
			successes++
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for err := range producer.Errors() {
			log.Println(err)
			errors++
		}
	}()

ProducerLoop:
	for {

		message := &sarama.ProducerMessage{
			Topic: "test",
			Value: sarama.StringEncoder("testing 123"),
		}

		select {
		case producer.Input() <- message:
			enqueued++

		case <-signals:
			producer.AsyncClose() // Trigger a shutdown of the producer.
			break ProducerLoop
		}
	}

	wg.Wait()

	log.Printf("Successfully produced: %d; errors: %d\n", successes, errors)

}
