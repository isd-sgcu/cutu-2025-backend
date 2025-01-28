package repository

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/storage/redis/v3"
)

type CacheRepository struct {
	redis *redis.Storage
}

func NewCacheRepository(redis *redis.Storage) *CacheRepository {
	return &CacheRepository{redis: redis}
}

// Set stores data in cache with expiration time
func (c *CacheRepository) Set(key string, value interface{}, expiration time.Duration) error {
	// Convert the value to JSON string
	jsonData, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %v", err)
	}

	// Store in Redis
	err = c.redis.Set(key, jsonData, expiration)
	if err != nil {
		return fmt.Errorf("failed to set cache: %v", err)
	}

	return nil
}

// Get retrieves data from cache
func (c *CacheRepository) Get(key string, result interface{}) error {
	// Get from Redis
	data, err := c.redis.Get(key)
	if err != nil {
		return fmt.Errorf("failed to get from cache: %v", err)
	}

	// If no data found
	if data == nil {
		return fmt.Errorf("cache miss for key: %s", key)
	}

	// Unmarshal the JSON data into the result interface
	err = json.Unmarshal(data, result)
	if err != nil {
		return fmt.Errorf("failed to unmarshal data: %v", err)
	}

	return nil
}

// Delete removes data from cache
func (c *CacheRepository) Delete(key string) error {
	err := c.redis.Delete(key)
	if err != nil {
		return fmt.Errorf("failed to delete from cache: %v", err)
	}

	return nil
}

// Clear removes all data from cache
func (c *CacheRepository) Clear() error {
	err := c.redis.Reset()
	if err != nil {
		return fmt.Errorf("failed to clear cache: %v", err)
	}

	return nil
}

// Has checks if a key exists in cache
func (c *CacheRepository) Has(key string) bool {
	data, err := c.redis.Get(key)
	return err == nil && data != nil
}

// SetMany stores multiple key-value pairs in cache
func (c *CacheRepository) SetMany(items map[string]interface{}, expiration time.Duration) error {
	for key, value := range items {
		if err := c.Set(key, value, expiration); err != nil {
			return fmt.Errorf("failed to set multiple items: %v", err)
		}
	}
	return nil
}

// GetMany retrieves multiple keys from cache
func (c *CacheRepository) GetMany(keys []string) (map[string][]byte, error) {
	results := make(map[string][]byte)

	for _, key := range keys {
		data, err := c.redis.Get(key)
		if err != nil {
			continue // Skip errors for individual keys
		}
		if data != nil {
			results[key] = data
		}
	}

	return results, nil
}
