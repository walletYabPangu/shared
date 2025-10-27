package models

import "time"

// ============================================
// Service & Instance Structs (اصلاح شده)
// ============================================

type ServiceInfo struct {
	ID              int                    `json:"id"`
	ServiceName     string                 `json:"service_name"`
	ServiceType     string                 `json:"service_type"`
	BaseURL         string                 `json:"base_url"`
	HealthCheckURL  *string                `json:"health_check_url,omitempty"` // اصلاح شد: پوینتر برای null
	Version         *string                `json:"version,omitempty"`          // اصلاح شد: پوینتر برای null
	Status          ServiceStatus          `json:"status"`                     // اصلاح شد: استفاده از Enum
	HealthStatus    HealthStatus           `json:"health_status"`              // اصلاح شد: استفاده از Enum
	Weight          int                    `json:"weight"`
	MaxConnections  int                    `json:"max_connections"`
	CircuitBreaker  CircuitBreakerConfig   `json:"circuit_breaker"`
	Tags            map[string]interface{} `json:"tags"`
	LastHealthCheck *time.Time             `json:"last_health_check,omitempty"` // اصلاح شد: پوینتر برای null
}

type CircuitBreakerConfig struct {
	Enabled   bool `json:"enabled"`
	Threshold int  `json:"threshold"`
	Timeout   int  `json:"timeout_seconds"`
}

type InstanceInfo struct {
	ID              int            `json:"id"`
	InstanceID      string         `json:"instance_id"`
	ServiceID       int            `json:"service_id"`
	BaseURL         string         `json:"base_url"`
	Status          InstanceStatus `json:"status"`                         // اصلاح شد: استفاده از Enum
	LastHeartbeat   *time.Time     `json:"last_heartbeat,omitempty"`       // اصلاح شد: پوینتر برای null
	CPUUsage        *float64       `json:"cpu_usage,omitempty"`            // اصلاح شد: پوینتر برای null
	MemoryUsage     *float64       `json:"memory_usage,omitempty"`         // اصلاح شد: پوینتر برای null
	AvgResponseTime *int           `json:"avg_response_time_ms,omitempty"` // اصلاح شد: پوینتر برای null
}
