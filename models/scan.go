package models

import "time"

type ScanSummary struct {
	TotalScanned int64            `json:"total_scanned"`
	PerChain     []ScanChainStats `json:"per_chain"`
}

type ScanChainStats struct {
	Chain   string `json:"chain"`
	Scanned int64  `json:"scanned"`
	Hits    int64  `json:"hits"`
}

type ScanItem struct {
	Chain       string    `json:"chain"`
	Address     string    `json:"address"`
	Balance     string    `json:"balance"`
	Block       int64     `json:"block"`
	Fingerprint string    `json:"fingerprint"`
	CreatedAt   time.Time `json:"created_at"`
}
