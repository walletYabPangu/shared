// shared/pkg/metrics/metrics.go
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	HTTPRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: []float64{0.001, 0.01, 0.1, 0.5, 1, 2, 5},
		},
		[]string{"service", "method", "path", "status"},
	)

	DBQueryDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "db_query_duration_seconds",
			Help:    "Duration of database queries",
			Buckets: []float64{0.001, 0.01, 0.1, 0.5, 1},
		},
		[]string{"service", "query"},
	)

	RedisOperationDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "redis_operation_duration_seconds",
			Help: "Duration of Redis operations",
		},
		[]string{"service", "operation"},
	)

	ActiveUsers = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "active_users",
			Help: "Number of active users",
		},
		[]string{"timeframe"},
	)

	TasksCompleted = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tasks_completed_total",
			Help: "Total number of tasks completed",
		},
		[]string{"task_type"},
	)

	ScanCreditsBalance = promauto.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "scan_credits_balance",
			Help:    "Distribution of user scan credit balances",
			Buckets: []float64{0, 10, 50, 100, 500, 1000, 5000},
		},
	)
)
