package models

import (
	"time"
	"project/shared/types"
)

type Task struct {
	ID          string             `json:"id"`
	Scope       string             `json:"scope"` // permanent, daily, challenge
	Type        string             `json:"type"`  // channel, twitter, story, referral
	Title       string             `json:"title"`
	Description string             `json:"description"`
	IconURL     string             `json:"icon_url"`
	RewardType  types.RewardType   `json:"reward_type"`
	RewardValue int64              `json:"reward_value"`
	IsActive    bool               `json:"is_active"`
	Meta        map[string]any     `json:"meta"`
	SortOrder   int                `json:"sort_order"`
}

type TaskTarget struct {
	ID         string `json:"id"`
	TaskID     string `json:"task_id"`
	TargetKind string `json:"target_kind"` // channel, post, url
	TargetID   string `json:"target_id"`
}

type UserTask struct {
	ID         string        `json:"id"`
	UserID     string        `json:"user_id"`
	TaskID     string        `json:"task_id"`
	Status     types.Status  `json:"status"`
	ProofURL   string        `json:"proof_url"`
	VerifiedAt *time.Time    `json:"verified_at"`
	ClaimedAt  *time.Time    `json:"claimed_at"`
	Progress   map[string]any `json:"progress"`
}
