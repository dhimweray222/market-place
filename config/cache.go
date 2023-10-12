package config

import (

	"github.com/go-redis/redis/v8"
	"log"
	"github.com/joho/godotenv"
	"os"
)

func init() {
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
}

func CreateRedisClient() *redis.Client {
		RedisAddr := os.Getenv("RedisAddr")
    client := redis.NewClient(&redis.Options{
        Addr:     RedisAddr, // Redis server address
        Password: "",              // No password by default
        DB:       0,               // Default database
    })

    return client
}


