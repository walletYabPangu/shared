package models

import "time"

type User struct {
	ID          string    `json:"id"`
	TgID        int64     `json:"tg_id"`
	Username    string    `json:"username"`
	JoinedAt    time.Time `json:"joined_at"`
	IsInChannel bool      `json:"is_in_channel"`
	WalletAddr  string    `json:"wallet_addr"`
}

type UserCounters struct {
	UserID         string `json:"user_id"`
	TotalFish      int64  `json:"total_fish"`
	TotalScans     int64  `json:"total_scans"`
	TotalSkins     int64  `json:"total_skins"`
	TotalOrders    int64  `json:"total_orders"`
	TotalReferrals int64  `json:"total_referrals"`
	TotalBoosts    int64  `json:"total_boosts"`
}
