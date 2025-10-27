package models

import "time"

type ServiceInfo struct {
	ID              int                    `json:"id"`
	ServiceName     string                 `json:"service_name"`
	ServiceType     string                 `json:"service_type"`
	BaseURL         string                 `json:"base_url"`
	HealthCheckURL  string                 `json:"health_check_url"`
	Version         string                 `json:"version"`
	Status          string                 `json:"status"`
	HealthStatus    string                 `json:"health_status"`
	Weight          int                    `json:"weight"`
	MaxConnections  int                    `json:"max_connections"`
	CircuitBreaker  CircuitBreakerConfig   `json:"circuit_breaker"`
	Tags            map[string]interface{} `json:"tags"`
	LastHealthCheck time.Time              `json:"last_health_check"`
}

type CircuitBreakerConfig struct {
	Enabled   bool `json:"enabled"`
	Threshold int  `json:"threshold"`
	Timeout   int  `json:"timeout_seconds"`
}

type InstanceInfo struct {
	ID              int       `json:"id"`
	InstanceID      string    `json:"instance_id"`
	ServiceID       int       `json:"service_id"`
	BaseURL         string    `json:"base_url"`
	Status          string    `json:"status"`
	LastHeartbeat   time.Time `json:"last_heartbeat"`
	CPUUsage        float64   `json:"cpu_usage"`
	MemoryUsage     float64   `json:"memory_usage"`
	AvgResponseTime int       `json:"avg_response_time_ms"`
}
