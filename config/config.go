package config

import (
	"log"

	"github.com/isd-sgcu/cutu2025-backend/utils"
	"github.com/joho/godotenv"
)

type Config struct {
	DBHost             string
	DBPort             string
	DBUser             string
	DBPassword         string
	DBName             string
	AWSRegion          string
	AWSAccessKeyID     string
	AWSSecretAccessKey string
	S3BucketName       string
	RedisHost          string
	RedisPort          string
	RedisPassword      string
}

// LoadConfig loads environment variables from .env and returns a Config struct
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return &Config{
		DBHost:             utils.GetEnv("DB_HOST", "localhost"),
		DBPort:             utils.GetEnv("DB_PORT", "5432"),
		DBUser:             utils.GetEnv("DB_USER", "postgres"),
		DBPassword:         utils.GetEnv("DB_PASSWORD", ""),
		DBName:             utils.GetEnv("DB_NAME", "postgres"),
		AWSRegion:          utils.GetEnv("AWS_REGION", "us-east-1"),
		AWSAccessKeyID:     utils.GetEnv("AWS_ACCESS_KEY_ID", ""),
		AWSSecretAccessKey: utils.GetEnv("AWS_SECRET_ACCESS_KEY", ""),
		S3BucketName:       utils.GetEnv("S3_BUCKET_NAME", ""),
		RedisHost:          utils.GetEnv("REDIS_HOST", "localhost"),
		RedisPort:          utils.GetEnv("REDIS_PORT", "6379"),
		RedisPassword:      utils.GetEnv("REDIS_PASSWORD", ""),
	}
}
