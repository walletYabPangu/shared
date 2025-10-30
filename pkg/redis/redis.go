// internal/pkg/redis/redis.go
package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/walletYabPangu/shared/config"
	"time"
)

type Client struct {
	*redis.Client
}

func New(cfg config.RedisConfig) (*Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("redis ping failed: %w", err)
	}

	return &Client{Client: client}, nil
}

// Helper methods
func (c *Client) SetJSON(ctx context.Context, key string, value interface{}, exp time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.Set(ctx, key, data, exp).Err()
}

func (c *Client) GetJSON(ctx context.Context, key string, dest interface{}) error {
	data, err := c.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, dest)
}

func (c *Client) SetNXWithRetry(ctx context.Context, key string, value interface{}, exp time.Duration, maxRetries int) (bool, error) {
	for i := 0; i < maxRetries; i++ {
		set, err := c.SetNX(ctx, key, value, exp).Result()
		if err != nil {
			if i == maxRetries-1 {
				return false, err
			}
			time.Sleep(time.Millisecond * 100 * time.Duration(i+1))
			continue
		}
		return set, nil
	}
	return false, nil
}

// Lua script for atomic increment with max check
const luaIncrWithMax = `
    local current = tonumber(redis.call('GET', KEYS[1]) or 0)
    local max = tonumber(ARGV[2])
    if current < max then
        return redis.call('INCR', KEYS[1])
    else
        return -1
    end
`

func (c *Client) IncrWithMax(ctx context.Context, key string, max int64, exp time.Duration) (int64, error) {
	result, err := c.Eval(ctx, luaIncrWithMax, []string{key}, 1, max).Int64()
	if err != nil {
		return 0, err
	}

	if result == -1 {
		return -1, fmt.Errorf("max limit reached")
	}

	c.Expire(ctx, key, exp)
	return result, nil
}
