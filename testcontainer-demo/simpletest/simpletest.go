package simpletest

import (
	"errors"
	"log"

	"github.com/go-redis/redis"
)

type Cache struct {
	client *redis.Client
}

func NewCache(endpoint string) *Cache {
	return &Cache{
		client: getRedisClient(endpoint),
	}
}

func getRedisClient(address string) *redis.Client {
	// Create a new Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     address, // Redis server address
		Password: "",      // Redis password (leave empty if no password is set)
		DB:       0,       // Redis database number
	})

	// Ping the Redis server to check if the connection is successful
	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("error on initializing redis %+v", err)
	}
	log.Println(pong)

	return client
}

func (cache *Cache) SetCache(key, value string) {
	cache.client.Set(key, value, 0)
}

func (cache *Cache) GetCache(key string) (string, error) {
	val, err := cache.client.Get(key).Result()
	switch {
	case errors.Is(err, redis.Nil):
		return "", nil
	case err != nil:
		return "", err
	}

	return val, nil
}
