// pkg/cache/cache.go
package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	client *redis.Client
}

func New(client *redis.Client) *Cache {
	return &Cache{client: client}
}

// Get with callback pattern (Cache-Aside)
func (c *Cache) GetOrSet(
	ctx context.Context,
	key string,
	ttl time.Duration,
	fetchFunc func() (interface{}, error),
	dest interface{},
) error {
	// Try to get from cache
	data, err := c.client.Get(ctx, key).Bytes()
	if err == nil {
		// Cache hit
		return json.Unmarshal(data, dest)
	}

	if err != redis.Nil {
		// Redis error (log but continue)
		fmt.Printf("Redis get error: %v\n", err)
	}

	// Cache miss - fetch from source
	result, err := fetchFunc()
	if err != nil {
		return err
	}

	// Unmarshal into dest
	resultBytes, _ := json.Marshal(result)
	if err := json.Unmarshal(resultBytes, dest); err != nil {
		return err
	}

	// Store in cache asynchronously
	go func() {
		data, _ := json.Marshal(result)
		c.client.Set(context.Background(), key, data, ttl).Err()
	}()

	return nil
}

// Set with immediate return (Write-Through)
func (c *Cache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.client.Set(ctx, key, data, ttl).Err()
}

// Delete from cache
func (c *Cache) Delete(ctx context.Context, keys ...string) error {
	if len(keys) == 0 {
		return nil
	}
	return c.client.Del(ctx, keys...).Err()
}

// Multi-key get
func (c *Cache) MGet(ctx context.Context, keys []string, dest interface{}) error {
	if len(keys) == 0 {
		return nil
	}

	results, err := c.client.MGet(ctx, keys...).Result()
	if err != nil {
		return err
	}

	// Filter non-nil results
	validResults := make([]interface{}, 0, len(results))
	for _, r := range results {
		if r != nil {
			validResults = append(validResults, r)
		}
	}

	if len(validResults) == 0 {
		return redis.Nil
	}

	// Unmarshal into dest
	data, _ := json.Marshal(validResults)
	return json.Unmarshal(data, dest)
}

// Hash operations with cache-aside
func (c *Cache) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return c.client.HGetAll(ctx, key).Result()
}

func (c *Cache) HSet(ctx context.Context, key string, values ...interface{}) error {
	return c.client.HSet(ctx, key, values...).Err()
}

func (c *Cache) HMSet(ctx context.Context, key string, fields map[string]interface{}) error {
	return c.client.HSet(ctx, key, fields).Err()
}

// Atomic operations
func (c *Cache) Incr(ctx context.Context, key string) (int64, error) {
	return c.client.Incr(ctx, key).Result()
}

func (c *Cache) IncrBy(ctx context.Context, key string, value int64) (int64, error) {
	return c.client.IncrBy(ctx, key, value).Result()
}

func (c *Cache) Decr(ctx context.Context, key string) (int64, error) {
	return c.client.Decr(ctx, key).Result()
}

func (c *Cache) DecrBy(ctx context.Context, key string, value int64) (int64, error) {
	return c.client.DecrBy(ctx, key, value).Result()
}

// Lock for distributed operations
func (c *Cache) Lock(ctx context.Context, key string, ttl time.Duration) (bool, error) {
	return c.client.SetNX(ctx, "lock:"+key, "1", ttl).Result()
}

func (c *Cache) Unlock(ctx context.Context, key string) error {
	return c.client.Del(ctx, "lock:"+key).Err()
}

// Pipeline for bulk operations
func (c *Cache) Pipeline(ctx context.Context, fn func(redis.Pipeliner) error) error {
	pipe := c.client.Pipeline()
	if err := fn(pipe); err != nil {
		return err
	}
	_, err := pipe.Exec(ctx)
	return err
}
