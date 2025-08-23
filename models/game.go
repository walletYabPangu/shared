package models

import "time"

type FishType struct {
	ID         string `json:"id"`
	Code       string `json:"code"`
	Title      string `json:"title"`
	ScanReward int64  `json:"scan_reward"`
	IsActive   bool   `json:"is_active"`
}

type FishCapture struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	TypeCode  string    `json:"type_code"`
	Quantity  int64     `json:"quantity"`
	RoundID   string    `json:"round_id"`
	CreatedAt time.Time `json:"created_at"`
}
