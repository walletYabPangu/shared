package models

import "time"

type ServiceRoute struct {
	// کلید سرویس (مثلاً "auth", "game", "user")
	ServiceKey string `gorm:"primaryKey"`

	// آدرس سرویس مقصد (مثلاً "http://localhost:8081")
	UpstreamURL string `gorm:"not null"`

	// GORM به صورت خودکار این دو فیلد را برای ردیابی زمان
	// ایجاد و آخرین به‌روزرسانی مدیریت می‌کند.
	CreatedAt time.Time
	UpdatedAt time.Time
}
