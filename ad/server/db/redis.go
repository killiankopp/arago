package db

import (
	"github.com/go-redis/redis/v8"
	"github.com/killiankopp/arago/ad/config"
	"golang.org/x/net/context"
	"log"
)

func ConnectToRedis() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: config.RedisURI,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func DisconnectFromRedis(client *redis.Client) {
	if err := client.Close(); err != nil {
		log.Fatalf("Failed to close Redis client: %v", err)
	}
}
