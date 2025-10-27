package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// ============================================
// ENUMS (به عنوان ثابت‌های Go)
// ============================================

type UserStatus string

const (
	UserStatusActive    UserStatus = "active"
	UserStatusBanned    UserStatus = "banned"
	UserStatusSuspended UserStatus = "suspended"
	UserStatusDeleted   UserStatus = "deleted"
)

type PaymentMethod string

const (
	PaymentMethodTON   PaymentMethod = "ton"
	PaymentMethodStars PaymentMethod = "stars"
	PaymentMethodFree  PaymentMethod = "free"
)

type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusConfirming OrderStatus = "confirming"
	OrderStatusConfirmed  OrderStatus = "confirmed"
	OrderStatusFailed     OrderStatus = "failed"
	OrderStatusRefunded   OrderStatus = "refunded"
	OrderStatusCancelled  OrderStatus = "cancelled"
)

type ServiceStatus string

const (
	ServiceStatusActive      ServiceStatus = "active"
	ServiceStatusInactive    ServiceStatus = "inactive"
	ServiceStatusMaintenance ServiceStatus = "maintenance"
	ServiceStatusDegraded    ServiceStatus = "degraded"
)

type HealthStatus string

const (
	HealthStatusHealthy   HealthStatus = "healthy"
	HealthStatusUnhealthy HealthStatus = "unhealthy"
	HealthStatusDegraded  HealthStatus = "degraded"
	HealthStatusUnknown   HealthStatus = "unknown"
)

type InstanceStatus string

const (
	InstanceStatusStarting  InstanceStatus = "starting"
	InstanceStatusHealthy   InstanceStatus = "healthy"
	InstanceStatusUnhealthy InstanceStatus = "unhealthy"
	InstanceStatusStopping  InstanceStatus = "stopping"
	InstanceStatusStopped   InstanceStatus = "stopped"
)

// ============================================
// USERS & PROFILES
// ============================================

type User struct {
	ID                   uint64  `gorm:"primarykey"`
	TelegramID           int64   `gorm:"not null;unique"`
	Username             *string `gorm:"type:varchar(255)"`
	FirstName            *string `gorm:"type:varchar(255)"`
	LastName             *string `gorm:"type:varchar(255)"`
	LanguageCode         string  `gorm:"type:varchar(10);default:'en'"`
	ProfilePhotoURL      *string `gorm:"type:text"`
	ProfilePhotoCachedAt *time.Time
	Bio                  *string `gorm:"type:text"`
	WalletAddr           *string `gorm:"type:varchar(255);index:idx_users_wallet_addr"`
	WalletConnectedAt    *time.Time
	WalletType           *string    `gorm:"type:varchar(20)"`
	WalletSignature      *string    `gorm:"type:text"`
	Status               UserStatus `gorm:"type:user_status;default:'active';index:idx_users_status"`
	IsInChannel          bool       `gorm:"default:false"`
	IsPremium            bool       `gorm:"default:false"`
	IsVerified           bool       `gorm:"default:false"`
	ReferralCode         *string    `gorm:"type:varchar(50);unique;index:idx_users_referral_code"`
	ReferredByCode       *string    `gorm:"type:varchar(50)"`
	ReferredByUserID     *uint64    // Foreign key
	ReferredByUser       *User      `gorm:"foreignKey:ReferredByUserID"`
	LastIP               *string    `gorm:"type:inet"`
	LastUserAgent        *string    `gorm:"type:text"`
	LoginCount           int        `gorm:"default:0"`
	JoinedAt             time.Time  `gorm:"not null;default:now();index:idx_users_joined_at"`
	LastActiveAt         time.Time  `gorm:"not null;default:now();index:idx_users_last_active"`
	BannedAt             *time.Time
	BannedReason         *string        `gorm:"type:text"`
	CreatedAt            time.Time      `gorm:"not null;default:now()"`
	UpdatedAt            time.Time      `gorm:"not null;default:now()"`
	DeletedAt            gorm.DeletedAt `gorm:"index"`
}

func (User) TableName() string { return "users" }

// ============================================
// COUNTERS & STATS
// ============================================

type UserCounter struct {
	UserID                   uint64          `gorm:"primarykey"` // 1-to-1 relationship
	User                     User            `gorm:"foreignKey:UserID"`
	TotalFishCaptured        int             `gorm:"default:0"`
	TotalRoundsPlayed        int             `gorm:"default:0"`
	TotalGameTimeSeconds     int64           `gorm:"default:0"`
	TotalScans               int             `gorm:"default:0"`
	TotalWalletsFound        int             `gorm:"default:0"`
	TotalScanCreditsEarned   int             `gorm:"default:0"`
	TotalScanCreditsSpent    int             `gorm:"default:0"`
	TotalTasksCompleted      int             `gorm:"default:0"`
	TotalDailyClaimed        int             `gorm:"default:0"`
	CurrentStreak            int             `gorm:"default:0"`
	MaxStreak                int             `gorm:"default:0"`
	TotalOrders              int             `gorm:"default:0"`
	TotalSpentTON            decimal.Decimal `gorm:"type:numeric(38,9);default:0"`
	TotalSpentStars          int             `gorm:"default:0"`
	TotalSkinsOwned          int             `gorm:"default:0"`
	TotalReferrals           int             `gorm:"default:0"`
	TotalChallengesCompleted int             `gorm:"default:0"`
	TotalBoostsPurchased     int             `gorm:"default:0"`
	TotalBoostsUsed          int             `gorm:"default:0"`
	UpdatedAt                time.Time       `gorm:"not null;default:now()"`
}

func (UserCounter) TableName() string { return "user_counters" }

type GlobalCounter struct {
	ID                     uint            `gorm:"primarykey;default:1"`
	TotalUsers             int64           `gorm:"default:0"`
	TotalActiveUsers24h    int64           `gorm:"default:0"`
	TotalActiveUsers7d     int64           `gorm:"default:0"`
	TotalPremiumUsers      int64           `gorm:"default:0"`
	TotalFishCaptured      int64           `gorm:"default:0"`
	TotalRoundsPlayed      int64           `gorm:"default:0"`
	TotalScans             int64           `gorm:"default:0"`
	TotalWalletsFound      int64           `gorm:"default:0"`
	TotalScanCreditsIssued int64           `gorm:"default:0"`
	TotalTasksCompleted    int64           `gorm:"default:0"`
	TotalOrders            int64           `gorm:"default:0"`
	TotalRevenueTON        decimal.Decimal `gorm:"type:numeric(38,9);default:0"`
	TotalRevenueStars      int64           `gorm:"default:0"`
	TotalReferrals         int64           `gorm:"default:0"`
	UpdatedAt              time.Time       `gorm:"not null;default:now()"`
}

func (GlobalCounter) TableName() string { return "global_counters" }

type DailyStat struct {
	StatDate        time.Time       `gorm:"primarykey;type:date;index:idx_daily_stats_date"`
	NewUsers        int             `gorm:"default:0"`
	ActiveUsers     int             `gorm:"default:0"`
	RoundsPlayed    int             `gorm:"default:0"`
	FishCaptured    int             `gorm:"default:0"`
	ScansCompleted  int             `gorm:"default:0"`
	WalletsFound    int             `gorm:"default:0"`
	TasksCompleted  int             `gorm:"default:0"`
	OrdersCreated   int             `gorm:"default:0"`
	OrdersConfirmed int             `gorm:"default:0"`
	RevenueTON      decimal.Decimal `gorm:"type:numeric(38,9);default:0"`
	RevenueStars    int             `gorm:"default:0"`
	CreatedAt       time.Time       `gorm:"not null;default:now()"`
}

func (DailyStat) TableName() string { return "daily_stats" }

// ============================================
// GAME SYSTEM
// ============================================

type FishTypeCfg struct {
	ID           uint            `gorm:"primarykey"`
	Code         string          `gorm:"type:varchar(20);not null;unique"`
	Title        string          `gorm:"type:varchar(100);not null"`
	Description  *string         `gorm:"type:text"`
	ScanReward   int             `gorm:"not null"`
	Rarity       string          `gorm:"type:varchar(20);default:'common'"`
	Probability  decimal.Decimal `gorm:"type:decimal(5,4)"`
	IconURL      *string         `gorm:"type:text"`
	AnimationURL *string         `gorm:"type:text"`
	IsActive     bool            `gorm:"default:true"`
	SortOrder    int             `gorm:"default:0"`
	CreatedAt    time.Time       `gorm:"not null;default:now()"`
	UpdatedAt    time.Time       `gorm:"not null;default:now()"`
	DeletedAt    gorm.DeletedAt  `gorm:"index"`
}

func (FishTypeCfg) TableName() string { return "fish_types_cfg" }

type GameConfig struct {
	ID                   uint      `gorm:"primarykey;default:1"`
	PlaysPerDay          int       `gorm:"default:3"`
	RoundDurationSeconds int       `gorm:"default:60"`
	ResetTime            string    `gorm:"type:time;default:'00:00:00'"` // GORM handles time.Time or string
	ResetTimezone        string    `gorm:"type:varchar(50);default:'UTC'"`
	MinFishPerRound      int       `gorm:"default:3"`
	MaxFishPerRound      int       `gorm:"default:8"`
	UpdatedAt            time.Time `gorm:"not null;default:now()"`
}

func (GameConfig) TableName() string { return "game_config" }

type FishCapture struct {
	ID                  uint64  `gorm:"primarykey"`
	UserID              uint64  `gorm:"not null;index:idx_captures_user_time"`
	User                User    `gorm:"foreignKey:UserID"`
	FishType            string  `gorm:"type:varchar(20);not null;index:idx_captures_fish_type"`
	Quantity            int     `gorm:"not null"`
	ScanRewardEarned    int     `gorm:"default:0;not null"`
	RoundID             *string `gorm:"type:varchar(100);index:idx_captures_round"`
	GameDurationSeconds *int
	CreatedAt           time.Time `gorm:"not null;default:now()"`
}

func (FishCapture) TableName() string { return "fish_captures" }

type UserDailyGame struct {
	StatDate      time.Time `gorm:"primaryKey;type:date;index:idx_daily_game_date"`
	UserID        uint64    `gorm:"primaryKey;index:idx_daily_game_user"`
	User          User      `gorm:"foreignKey:UserID"`
	PlaysUsed     int       `gorm:"not null;default:0"`
	PlaysBoosted  int       `gorm:"default:0"`
	FishCaught    int       `gorm:"default:0"`
	CreditsEarned int       `gorm:"default:0"`
	FirstPlayAt   *time.Time
	LastPlayAt    *time.Time
	UpdatedAt     time.Time `gorm:"not null;default:now()"`
}

func (UserDailyGame) TableName() string { return "user_daily_game" }

// ============================================
// BOOST SYSTEM
// ============================================

type Boost struct {
	ID         uint64           `gorm:"primarykey"`
	UserID     uint64           `gorm:"not null;index:idx_boosts_user_status;index:idx_boosts_active"`
	User       User             `gorm:"foreignKey:UserID"`
	BoostType  string           `gorm:"type:varchar(20);not null"`
	StartsAt   time.Time        `gorm:"not null;index:idx_boosts_active"`
	EndsAt     time.Time        `gorm:"not null;index:idx_boosts_user_status;index:idx_boosts_active"`
	PaidWith   PaymentMethod    `gorm:"type:payment_method;default:'free'"`
	PriceTON   *decimal.Decimal `gorm:"type:numeric(38,9)"`
	PriceStars *int
	Status     string `gorm:"type:varchar(20);default:'active';index:idx_boosts_user_status"`
	TimesUsed  int    `gorm:"default:0"`
	MaxUses    *int
	Metadata   datatypes.JSON
	CreatedAt  time.Time `gorm:"not null;default:now()"`
	UpdatedAt  time.Time `gorm:"not null;default:now()"`
}

func (Boost) TableName() string { return "boosts" }

// ============================================
// SCAN CREDITS & WALLET
// ============================================

type UserScanWallet struct {
	UserID            uint64 `gorm:"primarykey"`
	User              User   `gorm:"foreignKey:UserID"`
	Balance           int    `gorm:"not null;default:0;index:idx_wallet_balance"`
	LifetimeEarned    int    `gorm:"not null;default:0"`
	LifetimeSpent     int    `gorm:"not null;default:0"`
	LastTransactionAt *time.Time
	DailyEarnLimit    int        `gorm:"default:1000"`
	DailyEarnedToday  int        `gorm:"default:0"`
	DailyLimitResetAt *time.Time `gorm:"type:date"`
	CreatedAt         time.Time  `gorm:"not null;default:now()"`
	UpdatedAt         time.Time  `gorm:"not null;default:now()"`
}

func (UserScanWallet) TableName() string { return "user_scan_wallet" }

type UserScanLedger struct {
	ID            uint64  `gorm:"primarykey"`
	UserID        uint64  `gorm:"not null;index:idx_ledger_user_time"`
	User          User    `gorm:"foreignKey:UserID"`
	Delta         int     `gorm:"not null"`
	BalanceBefore int     `gorm:"not null"`
	BalanceAfter  int     `gorm:"not null"`
	Reason        string  `gorm:"type:varchar(50);not null;index:idx_ledger_reason"`
	RefType       *string `gorm:"type:varchar(50)"`
	RefID         *string `gorm:"type:varchar(100)"`
	Metadata      datatypes.JSON
	CreatedAt     time.Time `gorm:"not null;default:now()"`
}

func (UserScanLedger) TableName() string { return "user_scan_ledger" }

// ============================================
// SCAN SESSIONS & RESULTS
// ============================================

type ScanSession struct {
	ID                string  `gorm:"primarykey;type:varchar(50)"` // همانطور که می‌خواستید، کلید اصلی ساده
	UserID            uint64  `gorm:"not null;index:idx_sessions_user_status"`
	User              User    `gorm:"foreignKey:UserID"`
	RoundID           *string `gorm:"type:varchar(100)"`
	TotalScanned      int     `gorm:"not null;default:0"`
	TotalFound        int     `gorm:"not null;default:0"`
	ChainsScanned     datatypes.JSON
	CreditsSpent      int              `gorm:"not null;default:0"`
	Status            string           `gorm:"type:varchar(20);not null;default:'pending';index:idx_sessions_user_status;index:idx_sessions_flagged"`
	SpotVerified      bool             `gorm:"default:false;index:idx_sessions_verification"`
	VerificationScore *decimal.Decimal `gorm:"type:decimal(3,2);index:idx_sessions_verification"`
	VerificationNotes *string          `gorm:"type:text"`
	ClientVersion     *string          `gorm:"type:varchar(50)"`
	ClientFingerprint *string          `gorm:"type:varchar(64)"`
	ClientIP          *string          `gorm:"type:inet"`
	StartedAt         time.Time        `gorm:"not null;default:now()"`
	CompletedAt       *time.Time
	DurationSeconds   *int
	CreatedAt         time.Time `gorm:"not null;default:now()"`
	UpdatedAt         time.Time `gorm:"not null;default:now()"`
}

func (ScanSession) TableName() string { return "scan_sessions" }

type ScanResult struct {
	ID                 uint64           `gorm:"primarykey"`
	SessionID          string           `gorm:"type:varchar(50);not null;index:idx_results_session"`
	Session            ScanSession      `gorm:"foreignKey:SessionID"`
	Chain              string           `gorm:"type:varchar(10);not null;index:idx_results_chain_balance"`
	Address            string           `gorm:"type:varchar(255);not null"`
	BalanceBaseUnit    decimal.Decimal  `gorm:"type:numeric(78,0);not null;index:idx_results_chain_balance"`
	BalanceReadable    *decimal.Decimal `gorm:"type:decimal(30,18)"`
	UsdValue           *decimal.Decimal `gorm:"type:decimal(20,2)"`
	BlockNumber        *int64
	LastTransactionAt  *time.Time
	TransactionCount   *int
	MnemonicEncrypted  []byte
	EncryptionKeyID    *string `gorm:"type:varchar(50)"`
	WalletType         *string `gorm:"type:varchar(20)"`
	WalletAgeDays      *int
	Verified           bool `gorm:"default:false;index:idx_results_verified"`
	VerifiedAt         *time.Time
	VerificationMethod *string   `gorm:"type:varchar(20)"`
	CreatedAt          time.Time `gorm:"not null;default:now()"`
}

func (ScanResult) TableName() string { return "scan_results" }

// ============================================
// TASKS SYSTEM
// ============================================

type Task struct {
	ID                 uint    `gorm:"primarykey"`
	Scope              string  `gorm:"type:varchar(20);not null;index:idx_tasks_scope_active"`
	TaskType           string  `gorm:"type:varchar(20);not null"`
	Title              string  `gorm:"type:varchar(200);not null"`
	Description        *string `gorm:"type:text"`
	IconURL            *string `gorm:"type:text"`
	ImageURL           *string `gorm:"type:text"`
	RewardType         *string `gorm:"type:varchar(20)"`
	RewardValue        *int
	RewardMetadata     datatypes.JSON
	RequirementType    *string    `gorm:"type:varchar(50)"`
	RequirementValue   *string    `gorm:"type:varchar(500)"`
	RequirementCount   int        `gorm:"default:1"`
	IsActive           bool       `gorm:"default:true;index:idx_tasks_scope_active;index:idx_tasks_availability"`
	StartsAt           *time.Time `gorm:"index:idx_tasks_availability"`
	EndsAt             *time.Time `gorm:"index:idx_tasks_availability"`
	MaxCompletions     *int
	CurrentCompletions int     `gorm:"default:0"`
	SortOrder          int     `gorm:"default:0;index:idx_tasks_scope_active"`
	Category           *string `gorm:"type:varchar(50)"`
	Tags               datatypes.JSON
	CreatedAt          time.Time      `gorm:"not null;default:now()"`
	UpdatedAt          time.Time      `gorm:"not null;default:now()"`
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}

func (Task) TableName() string { return "tasks" }

type TaskTarget struct {
	ID                 uint      `gorm:"primarykey"`
	TaskID             uint      `gorm:"not null;uniqueIndex:idx_task_target_unique"`
	Task               Task      `gorm:"foreignKey:TaskID"`
	TargetType         string    `gorm:"type:varchar(20);not null;uniqueIndex:idx_task_target_unique"`
	TargetID           string    `gorm:"type:varchar(200);not null;uniqueIndex:idx_task_target_unique"`
	TargetURL          *string   `gorm:"type:text"`
	VerificationMethod *string   `gorm:"type:varchar(50)"`
	CreatedAt          time.Time `gorm:"not null;default:now()"`
}

func (TaskTarget) TableName() string { return "task_targets" }

type UserTask struct {
	ID               uint64 `gorm:"primarykey"`
	UserID           uint64 `gorm:"not null;uniqueIndex:idx_user_task_unique;index:idx_user_tasks_user_status"`
	User             User   `gorm:"foreignKey:UserID"`
	TaskID           uint   `gorm:"not null;uniqueIndex:idx_user_task_unique;index:idx_user_tasks_task"`
	Task             Task   `gorm:"foreignKey:TaskID"`
	Status           string `gorm:"type:varchar(20);not null;default:'pending';index:idx_user_tasks_user_status;index:idx_user_tasks_task;index:idx_user_tasks_pending"`
	ProgressCurrent  int    `gorm:"default:0"`
	ProgressRequired int    `gorm:"default:1"`
	ProgressMetadata datatypes.JSON
	ProofType        *string `gorm:"type:varchar(20)"`
	ProofURL         *string `gorm:"type:text"`
	ProofData        datatypes.JSON
	StartedAt        *time.Time
	VerifiedAt       *time.Time
	ClaimedAt        *time.Time
	FailedAt         *time.Time
	FailureReason    *string   `gorm:"type:text"`
	CreatedAt        time.Time `gorm:"not null;default:now()"`
	UpdatedAt        time.Time `gorm:"not null;default:now();index:idx_user_tasks_pending"`
}

func (UserTask) TableName() string { return "user_tasks" }

// ============================================
// DAILY STREAK SYSTEM
// ============================================

type UserDailyStreak struct {
	UserID               uint64 `gorm:"primarykey"`
	User                 User   `gorm:"foreignKey:UserID"`
	CurrentDay           int    `gorm:"not null;default:1"`
	StagesTotal          int    `gorm:"not null;default:7"`
	LastClaimedAt        *time.Time
	LastClaimedDay       *time.Time `gorm:"type:date"`
	TotalCyclesCompleted int        `gorm:"default:0"`
	CreatedAt            time.Time  `gorm:"not null;default:now()"`
	UpdatedAt            time.Time  `gorm:"not null;default:now()"`
}

func (UserDailyStreak) TableName() string { return "user_daily_streak" }

type DailyStreakStage struct {
	ID             uint   `gorm:"primarykey"`
	DayIndex       int    `gorm:"not null;unique"`
	TaskID         *uint  // Nullable Foreign key
	Task           *Task  `gorm:"foreignKey:TaskID"`
	RewardType     string `gorm:"type:varchar(20);not null"`
	RewardValue    int    `gorm:"not null"`
	RewardMetadata datatypes.JSON
	IconURL        *string   `gorm:"type:text"`
	Title          *string   `gorm:"type:varchar(100)"`
	IsActive       bool      `gorm:"default:true"`
	CreatedAt      time.Time `gorm:"not null;default:now()"`
	UpdatedAt      time.Time `gorm:"not null;default:now()"`
}

func (DailyStreakStage) TableName() string { return "daily_streak_stages" }

// ============================================
// REFERRAL SYSTEM
// ============================================

type ReferralCode struct {
	Code      string `gorm:"primarykey;type:varchar(50)"`
	OwnerID   uint64 `gorm:"not null;index:idx_referral_owner"`
	Owner     User   `gorm:"foreignKey:OwnerID"`
	IsActive  bool   `gorm:"default:true"`
	TotalUses int    `gorm:"default:0"`
	MaxUses   *int
	ExpiresAt *time.Time
	CreatedAt time.Time `gorm:"not null;default:now()"`
}

func (ReferralCode) TableName() string { return "referral_codes" }

type ReferralUse struct {
	ID           uint64       `gorm:"primarykey"`
	Code         string       `gorm:"type:varchar(50);not null;uniqueIndex:idx_referral_use_unique;index:idx_referral_uses_code"`
	ReferralCode ReferralCode `gorm:"foreignKey:Code"`
	RefereeID    uint64       `gorm:"not null;uniqueIndex:idx_referral_use_unique;index:idx_referral_uses_referee"`
	Referee      User         `gorm:"foreignKey:RefereeID"`
	RewardGiven  bool         `gorm:"default:false"`
	RewardType   *string      `gorm:"type:varchar(20)"`
	RewardValue  *int
	CreatedAt    time.Time `gorm:"not null;default:now()"`
}

func (ReferralUse) TableName() string { return "referral_uses" }

type ReferralReward struct {
	ID                uint   `gorm:"primarykey"`
	ReferralsRequired int    `gorm:"not null;unique"`
	RewardType        string `gorm:"type:varchar(20);not null"`
	RewardValue       int    `gorm:"not null"`
	RewardMetadata    datatypes.JSON
	Title             *string   `gorm:"type:varchar(100)"`
	Description       *string   `gorm:"type:text"`
	IsActive          bool      `gorm:"default:true"`
	CreatedAt         time.Time `gorm:"not null;default:now()"`
}

func (ReferralReward) TableName() string { return "referral_rewards" }

// ============================================
// CHALLENGE SYSTEM
// ============================================

type Challenge struct {
	ID               uint    `gorm:"primarykey"`
	Code             string  `gorm:"type:varchar(50);not null;unique"`
	Title            string  `gorm:"type:varchar(200);not null"`
	Description      *string `gorm:"type:text"`
	ChallengeType    string  `gorm:"type:varchar(20);not null"`
	RequirementType  *string `gorm:"type:varchar(50)"`
	RequirementValue *int
	RewardType       string `gorm:"type:varchar(20);not null"`
	RewardValue      int    `gorm:"not null"`
	RewardMetadata   datatypes.JSON
	HasLeaderboard   bool      `gorm:"default:false"`
	LeaderboardSize  int       `gorm:"default:100"`
	StartsAt         time.Time `gorm:"not null;index:idx_challenges_active"`
	EndsAt           time.Time `gorm:"not null;index:idx_challenges_active"`
	IsActive         bool      `gorm:"default:true;index:idx_challenges_active"`
	CreatedAt        time.Time `gorm:"not null;default:now()"`
	UpdatedAt        time.Time `gorm:"not null;default:now()"`
}

func (Challenge) TableName() string { return "challenges" }

type UserChallenge struct {
	ID               uint64    `gorm:"primarykey"`
	UserID           uint64    `gorm:"not null;uniqueIndex:idx_user_challenge_unique;index:idx_user_challenges_user"`
	User             User      `gorm:"foreignKey:UserID"`
	ChallengeID      uint      `gorm:"not null;uniqueIndex:idx_user_challenge_unique;index:idx_user_challenges_challenge_score"`
	Challenge        Challenge `gorm:"foreignKey:ChallengeID"`
	Status           string    `gorm:"type:varchar(20);not null;default:'active';index:idx_user_challenges_user"`
	ProgressCurrent  int       `gorm:"default:0"`
	ProgressRequired *int
	Score            int `gorm:"default:0;index:idx_user_challenges_challenge_score"`
	Rank             *int
	CompletedAt      *time.Time
	RewardedAt       *time.Time
	Metadata         datatypes.JSON
	CreatedAt        time.Time `gorm:"not null;default:now()"`
	UpdatedAt        time.Time `gorm:"not null;default:now()"`
}

func (UserChallenge) TableName() string { return "user_challenges" }

// ============================================
// SHOP & PAYMENTS
// ============================================

type Skin struct {
	ID           uint             `gorm:"primarykey"`
	Name         string           `gorm:"type:varchar(100);not null"`
	Description  *string          `gorm:"type:text"`
	Category     *string          `gorm:"type:varchar(50)"`
	SupplyTotal  int              `gorm:"not null"`
	SupplySold   int              `gorm:"not null;default:0"`
	PriceTON     *decimal.Decimal `gorm:"type:numeric(38,9)"`
	PriceStars   *int
	Rarity       string  `gorm:"type:varchar(20);default:'common'"`
	MediaURL     *string `gorm:"type:text"`
	ThumbnailURL *string `gorm:"type:text"`
	PreviewURLs  datatypes.JSON
	IsActive     bool           `gorm:"default:true;index:idx_skins_active"`
	IsFeatured   bool           `gorm:"default:false;index:idx_skins_featured"`
	SortOrder    int            `gorm:"default:0;index:idx_skins_active"`
	CreatedAt    time.Time      `gorm:"not null;default:now()"`
	UpdatedAt    time.Time      `gorm:"not null;default:now()"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (Skin) TableName() string { return "skins" }

type Order struct {
	ID                     uint64           `gorm:"primarykey"`
	UserID                 uint64           `gorm:"not null;index:idx_orders_user_status"`
	User                   User             `gorm:"foreignKey:UserID"`
	OrderType              string           `gorm:"type:varchar(20);not null"`
	SkinID                 *uint            // Nullable Foreign key
	Skin                   *Skin            `gorm:"foreignKey:SkinID"`
	Quantity               int              `gorm:"not null;default:1"`
	PaymentMethod          PaymentMethod    `gorm:"type:payment_method;not null"`
	AmountTON              *decimal.Decimal `gorm:"type:numeric(38,9)"`
	AmountStars            *int
	TxHash                 *string `gorm:"type:varchar(100);index:idx_orders_tx_hash"`
	TxLt                   *int64
	TxFromAddress          *string          `gorm:"type:varchar(255)"`
	TxToAddress            *string          `gorm:"type:varchar(255)"`
	TxValue                *decimal.Decimal `gorm:"type:numeric(38,9)"`
	TxConfirmedAt          *time.Time
	TxConfirmations        int         `gorm:"default:0"`
	TelegramPaymentID      *string     `gorm:"type:varchar(100);index:idx_orders_telegram_payment"`
	TelegramChargeID       *string     `gorm:"type:varchar(100)"`
	TelegramInvoicePayload *string     `gorm:"type:text"`
	Status                 OrderStatus `gorm:"type:order_status;default:'pending';index:idx_orders_user_status;index:idx_orders_status_created"`
	FulfilledAt            *time.Time
	ErrorCode              *string `gorm:"type:varchar(50)"`
	ErrorMessage           *string `gorm:"type:text"`
	Metadata               datatypes.JSON
	CreatedAt              time.Time `gorm:"not null;default:now();index:idx_orders_status_created"`
	UpdatedAt              time.Time `gorm:"not null;default:now()"`
}

func (Order) TableName() string { return "orders" }

type UserSkin struct {
	ID          uint64    `gorm:"primarykey"`
	UserID      uint64    `gorm:"not null;uniqueIndex:idx_user_skin_unique;index:idx_user_skins_user;index:idx_user_skins_equipped"`
	User        User      `gorm:"foreignKey:UserID"`
	SkinID      uint      `gorm:"not null;uniqueIndex:idx_user_skin_unique"`
	Skin        Skin      `gorm:"foreignKey:SkinID"`
	Quantity    int       `gorm:"not null;default:1"`
	IsEquipped  bool      `gorm:"default:false;index:idx_user_skins_equipped"`
	AcquiredVia *string   `gorm:"type:varchar(20)"`
	OrderID     *uint64   // Nullable Foreign key
	Order       *Order    `gorm:"foreignKey:OrderID"`
	AcquiredAt  time.Time `gorm:"not null;default:now()"`
}

func (UserSkin) TableName() string { return "user_skins" }

type PaymentWallet struct {
	ID               uint            `gorm:"primarykey"`
	WalletType       string          `gorm:"type:varchar(20);not null"`
	Address          string          `gorm:"type:varchar(255);not null;unique"`
	PublicKey        *string         `gorm:"type:text"`
	IsActive         bool            `gorm:"default:true"`
	IsPrimary        bool            `gorm:"default:false"`
	BalanceTON       decimal.Decimal `gorm:"type:numeric(38,9);default:0"`
	LastBalanceCheck *time.Time
	Metadata         datatypes.JSON
	CreatedAt        time.Time `gorm:"not null;default:now()"`
	UpdatedAt        time.Time `gorm:"not null;default:now()"`
}

func (PaymentWallet) TableName() string { return "payment_wallets" }

type PaymentWebhook struct {
	ID                uint64         `gorm:"primarykey"`
	Source            string         `gorm:"type:varchar(20);not null"`
	WebhookType       *string        `gorm:"type:varchar(50)"`
	Payload           datatypes.JSON `gorm:"not null"`
	Signature         *string        `gorm:"type:varchar(500)"`
	SignatureVerified *bool
	OrderID           *uint64 `gorm:"index:idx_webhooks_order"`
	Order             *Order  `gorm:"foreignKey:OrderID"`
	Processed         bool    `gorm:"default:false;index:idx_webhooks_processed"`
	ProcessedAt       *time.Time
	ErrorMessage      *string   `gorm:"type:text"`
	CreatedAt         time.Time `gorm:"not null;default:now();index:idx_webhooks_processed"`
}

func (PaymentWebhook) TableName() string { return "payment_webhooks" }

// ============================================
// SERVICE DISCOVERY & REGISTRY
// ============================================

type ServiceRegistry struct {
	ID                           uint          `gorm:"primarykey"`
	ServiceName                  string        `gorm:"type:varchar(50);not null;unique;index:idx_service_registry_name_status"`
	ServiceType                  string        `gorm:"type:varchar(20);not null"`
	BaseURL                      string        `gorm:"type:varchar(255);not null"`
	HealthCheckURL               *string       `gorm:"type:varchar(255)"`
	Version                      *string       `gorm:"type:varchar(20)"`
	Region                       string        `gorm:"type:varchar(50);default:'default'"`
	Environment                  string        `gorm:"type:varchar(20);default:'production'"`
	Status                       ServiceStatus `gorm:"type:service_status;default:'active';index:idx_service_registry_name_status"`
	HealthStatus                 HealthStatus  `gorm:"type:health_status;default:'unknown';index:idx_service_registry_health"`
	LastHealthCheck              *time.Time    `gorm:"index:idx_service_registry_health"`
	ConsecutiveFailures          int           `gorm:"default:0"`
	Weight                       int           `gorm:"default:100"`
	MaxConnections               int           `gorm:"default:100"`
	CurrentConnections           int           `gorm:"default:0"`
	RateLimitPerMinute           int           `gorm:"default:1000"`
	CircuitBreakerEnabled        bool          `gorm:"default:true"`
	CircuitBreakerThreshold      int           `gorm:"default:5"`
	CircuitBreakerTimeoutSeconds int           `gorm:"default:60"`
	Tags                         datatypes.JSON
	Config                       datatypes.JSON
	CreatedAt                    time.Time `gorm:"not null;default:now()"`
	UpdatedAt                    time.Time `gorm:"not null;default:now()"`
	CreatedBy                    *string   `gorm:"type:varchar(50)"`
}

func (ServiceRegistry) TableName() string { return "service_registry" }

type ServiceInstance struct {
	ID                uint             `gorm:"primarykey"`
	ServiceID         uint             `gorm:"not null;index:idx_instances_service_status"`
	Service           ServiceRegistry  `gorm:"foreignKey:ServiceID"`
	InstanceID        string           `gorm:"type:varchar(100);not null;unique"`
	Host              string           `gorm:"type:varchar(255);not null"`
	Port              int              `gorm:"not null"`
	BaseURL           string           `gorm:"type:varchar(255);not null"`
	ContainerID       *string          `gorm:"type:varchar(100)"`
	NodeName          *string          `gorm:"type:varchar(100)"`
	PodName           *string          `gorm:"type:varchar(100)"`
	Status            InstanceStatus   `gorm:"type:instance_status;default:'starting';index:idx_instances_service_status;index:idx_instances_heartbeat"`
	LastHeartbeat     *time.Time       `gorm:"index:idx_instances_heartbeat"`
	UptimeSeconds     int64            `gorm:"default:0"`
	CPUUsage          *decimal.Decimal `gorm:"type:decimal(5,2)"`
	MemoryUsage       *decimal.Decimal `gorm:"type:decimal(5,2)"`
	RequestCount      int64            `gorm:"default:0"`
	ErrorCount        int64            `gorm:"default:0"`
	AvgResponseTimeMs *int
	Version           *string   `gorm:"type:varchar(20)"`
	StartedAt         time.Time `gorm:"not null;default:now()"`
	Metadata          datatypes.JSON
	CreatedAt         time.Time `gorm:"not null;default:now()"`
	UpdatedAt         time.Time `gorm:"not null;default:now()"`
}

func (ServiceInstance) TableName() string { return "service_instances" }

type ServiceHealthHistory struct {
	ID             uint64 `gorm:"primarykey"`
	ServiceID      int    `gorm:"not null;index:idx_health_history_service"`
	InstanceID     *int
	CheckType      *string `gorm:"type:varchar(20)"`
	Status         *string `gorm:"type:varchar(20)"`
	ResponseTimeMs *int
	StatusCode     *int
	ErrorMessage   *string   `gorm:"type:text"`
	CheckedAt      time.Time `gorm:"not null;default:now()"`
}

func (ServiceHealthHistory) TableName() string { return "service_health_history" }

// ============================================
// AUDIT & SECURITY
// ============================================

type AuditLog struct {
	ID        uint64  `gorm:"primarykey"`
	Entity    string  `gorm:"type:varchar(50);not null;index:idx_audit_entity"`
	EntityID  *string `gorm:"type:varchar(100);index:idx_audit_entity"`
	Action    string  `gorm:"type:varchar(50);not null"`
	ActorType *string `gorm:"type:varchar(20)"`
	ActorID   *uint64 `gorm:"index:idx_audit_actor"`
	Changes   datatypes.JSON
	IPAddress *string   `gorm:"type:inet"`
	UserAgent *string   `gorm:"type:text"`
	Severity  string    `gorm:"type:varchar(20);default:'info';index:idx_audit_severity"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
}

func (AuditLog) TableName() string { return "audit_logs" }

type SecurityEvent struct {
	ID         uint64  `gorm:"primarykey"`
	EventType  string  `gorm:"type:varchar(50);not null"`
	UserID     *uint64 `gorm:"index:idx_security_events_user"`
	User       *User   `gorm:"foreignKey:UserID"`
	IPAddress  *string `gorm:"type:inet"`
	UserAgent  *string `gorm:"type:text"`
	Severity   string  `gorm:"type:varchar(20);default:'warning'"`
	Details    datatypes.JSON
	Resolved   bool `gorm:"default:false;index:idx_security_events_unresolved"`
	ResolvedAt *time.Time
	ResolvedBy *string   `gorm:"type:varchar(100)"`
	CreatedAt  time.Time `gorm:"not null;default:now()"`
}

func (SecurityEvent) TableName() string { return "security_events" }

// ============================================
// MONITORING & METRICS
// ============================================

type SystemMetric struct {
	ID          uint64           `gorm:"primarykey"`
	MetricName  string           `gorm:"type:varchar(100);not null;index:idx_metrics_name_time"`
	MetricValue *decimal.Decimal `gorm:"type:numeric"`
	Labels      datatypes.JSON
	RecordedAt  time.Time `gorm:"not null;default:now()"`
}

func (SystemMetric) TableName() string { return "system_metrics" }
