package models

import (
	"github.com/walletYabPangu/shared/types"
	"time"
)

type Boost struct {
	ID       string          `json:"id"`
	UserID   string          `json:"user_id"`
	Kind     types.BoostKind `json:"kind"`
	StartsAt time.Time       `json:"starts_at"`
	EndsAt   time.Time       `json:"ends_at"`
	Status   string          `json:"status"` // active, expired
}
