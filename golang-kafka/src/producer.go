package main

import (
  "fmt"
  "github.com/confluentinc/confluent-kafka-go/kafka"
  "sync"
)

var (
  broker  = "localhost:29092"
  groupId = "my-group"
  topic   = "topic-name"
)

func startProducer() {
  p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
  if err != nil {
    panic(err)
  }
  var wg sync.WaitGroup

  
  go func() {
    for e := range p.Events() {
      switch ev := e.(type) {
      case *kafka.Message:
        if ev.TopicPartition.Error != nil {
          fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
        } else {
          fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
          wg.Done()
        }
      }
    }
  }()
  
  for _, word := range []string{"message 1", "message 2", "message 3"} {
    p.Produce(&kafka.Message{
      TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
      Value:          []byte(word),
    }, nil)
    fmt.Println(word)
    wg.Add(1)
  }

  wg.Wait()
  p.Close()
}


func main() {
  startProducer()
}
