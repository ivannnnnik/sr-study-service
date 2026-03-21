package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"google.golang.org/genproto/googleapis/maps/fleetengine/delivery/v1"
)

type Producer struct{
	producer *kafka.Producer
	topic string
}

func NewProducer(brokers, topic string) (*Producer, error){
	p, err := kafka.NewProducer(
		&kafka.ConfigMap{
			"bootsrap.servers": brokers,
		},
	)

	if err != nil{
		return nil, fmt.Errorf("Kafka new producer error: %w", err)
	}

	return &Producer{
		producer: p,
		topic: topic,
	}, nil

}


func (p *Producer) Send(ctx context.Context, key string, event any) error{
	value, err := json.Marshal(event)
	
	if err != nil{
		return fmt.Errorf("Marshal event error: %w", err)
	}

	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic: &p.topic,
			Partition: kafka.PartitionAny,
		},
		Key: []byte(key),
		Value: value,
	}

	deliveryChan := make(chan kafka.Event, 1)
	if err := p.producer.Produce(msg, deliveryChan); err != nil{
		return fmt.Errorf("produce error: %w", err)
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil{
		return fmt.Errorf("delivery failed: %w", m.TopicPartition.Error)
	}

	return nil
}

func (p *Producer) Close(){
	p.producer.Flush(5000)
	p.producer.Close()
}
