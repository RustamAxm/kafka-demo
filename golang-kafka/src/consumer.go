package main

import (
  "fmt"
  "github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
  broker  = "localhost:29092"
  groupId = "my-group"
  topic   = "topic-name"
)

func startConsumer() {
  c, err := kafka.NewConsumer(&kafka.ConfigMap{
    "bootstrap.servers": broker,
    "group.id":          groupId,
    "auto.offset.reset": "earliest",
  })

  if err != nil {
    panic(err)
  }
  c.Subscribe(topic, nil)

  for {
    msg, err := c.ReadMessage(-1)
    if err == nil {
      fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
    } else {
      fmt.Printf("Consumer error: %v (%v)\n", err, msg)
      break
    }
  }

  c.Close()
}

func main() {
  startConsumer()
}
