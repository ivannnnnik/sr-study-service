package kafka

import (
	"fmt"
	
	"log/slog"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"

)

type Consumer struct{
	consumer *kafka.Consumer
	logger *slog.Logger
}

func NewConsumer(brokers, groupID, topic string, logger *slog.Logger)(*Consumer, error){
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bottstrap.servers": brokers,
		"group.id": groupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil{
		return  nil, fmt.Errorf("kafka new comsumer: %w", err)
	}

	if err := c.Subscribe(topic, nil); err != nil{
		c.Close()
		return nil, fmt.Errorf("Kafka subscribe: %w", err)
	}

	return &Consumer{
		consumer: c,
		logger: logger,
	}, nil
}

