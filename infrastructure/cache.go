package infrastructure

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/storage/redis/v3"
	"github.com/isd-sgcu/cutu2025-backend/config"
)

// ConnectRedis initializes and returns a Redis client
func ConnectRedis(cfg *config.Config) *redis.Storage {
	port, err := strconv.Atoi(cfg.RedisPort)
	if err != nil {
		log.Fatalf("Invalid Redis port: %v", err)
	}

	storage := redis.New(redis.Config{
		Host:     cfg.RedisHost,
		Port:     port,
		Password: cfg.RedisPassword,
		Database: 0,
	})

	fmt.Println("Connected to Redis successfully")
	return storage
}
