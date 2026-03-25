package kafka

import (
	"fmt"
	"context"
	"encoding/json"

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

func (c *Consumer) Run(ctx context.Context) {
	for {
		select {
		case <- ctx.Done():
			return
		default: 
			msg, err := c.consumer.ReadMessage(-1)
			if err != nil{
				c.logger.Error("read message: %w", err)
			}

			var event StudyEvent
			if err := json.Unmarshal(msg.Value, &event); err != nil{
				c.logger.Error("unmarshall event: %w", err)
			}

			c.logger.Info(
				"study event received", 
				"user_id", event.UserID,
				"question_id", event.QuestionID,
				"quality", event.Quality,
				"ease_factor", event.EaseFactor,
				"interval", event.Interval,
			)
		}
	}
}
