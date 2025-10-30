// internal/pkg/logger/logger.go
package logger

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger
}

func New(env string) *Logger {
	var config zap.Config

	if env == "production" {
		config = zap.NewProductionConfig()
		config.EncoderConfig.TimeKey = "timestamp"
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	} else {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	logger, _ := config.Build()
	return &Logger{logger.Sugar()}
}

func (l *Logger) WithContext(ctx context.Context) *Logger {
	// Extract request ID, user ID, etc. from context
	fields := []interface{}{}

	if reqID, ok := ctx.Value("request_id").(string); ok {
		fields = append(fields, "request_id", reqID)
	}

	if userID, ok := ctx.Value("user_id").(int64); ok {
		fields = append(fields, "user_id", userID)
	}

	return &Logger{l.SugaredLogger.With(fields...)}
}

func (l *Logger) WithService(serviceName string) *Logger {
	return &Logger{l.SugaredLogger.With("service", serviceName)}
}
