// consumer/message_consumer.go
package consumer

import (
	"context"
	"encoding/json"
	"log"
	"websocket-messaging/internal/models"
	"websocket-messaging/internal/redis"
)

type MessageConsumer struct {
	RedisClient redis.RedisClient
}

// NewMessageConsumer creates a new instance of MessageConsumer
func GetMessageConsumer(redisClient redis.RedisClient) *MessageConsumer {
	return &MessageConsumer{redisClient}
}

func (mc *MessageConsumer) ProcessMessage(ctx context.Context, message string) (bool, error) {
	log.Println("Processing message:", message)
	var messageData models.Message
	err := json.Unmarshal([]byte(message), &messageData)
	if err != nil {
		log.Println("Failed to unmarshal message:", err)
		return false, err
	}
	serverId := mc.RedisClient.GetUserServer(messageData.ReceiverID)
	go mc.RedisClient.Publish(serverId, message)
	return true, nil
}

func (mc *MessageConsumer) ProcessBulkMessage(ctx context.Context, messages []string) (bool, error) {
	log.Println("Processing bulk messages:", messages)
	return true, nil
}

func (mc *MessageConsumer) GetConsumerName() string {
	return "MessageConsumer"
}
