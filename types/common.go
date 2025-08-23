package types

import "time"

type RewardType string

const (
	RewardScan       RewardType = "scan"
	RewardPoints     RewardType = "points"
	RewardBoostDaily RewardType = "boost_daily"
	RewardFish       RewardType = "fish"
	RewardOther      RewardType = "other"
)

type Status string

const (
	StatusPending  Status = "pending"
	StatusVerified Status = "verified"
	StatusFailed   Status = "failed"
	StatusClaimed  Status = "claimed"
)

type BoostKind string

const (
	BoostDaily BoostKind = "daily"
	Boost24h   BoostKind = "24h"
)

type BaseModel struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
