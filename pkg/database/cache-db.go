// internal/pkg/database/cached_db.go
package database

import (
	"context"
	"time"

	"github.com/walletYabPangu/shared/pkg/cache"

	"gorm.io/gorm"
)

type CachedDB struct {
	DB    *gorm.DB
	Cache *cache.Cache
}

func NewCachedDB(db *gorm.DB, cache *cache.Cache) *CachedDB {
	return &CachedDB{
		DB:    db,
		Cache: cache,
	}
}

// Get single record with cache
func (cdb *CachedDB) GetWithCache(
	ctx context.Context,
	cacheKey string,
	ttl time.Duration,
	dest interface{},
	query func(*gorm.DB) *gorm.DB,
) error {
	return cdb.Cache.GetOrSet(ctx, cacheKey, ttl, func() (interface{}, error) {
		if err := query(cdb.DB.WithContext(ctx)).First(dest).Error; err != nil {
			return nil, err
		}
		return dest, nil
	}, dest)
}

// Update with cache invalidation
func (cdb *CachedDB) UpdateWithCache(
	ctx context.Context,
	cacheKeys []string,
	updateFunc func(*gorm.DB) error,
) error {
	// Perform DB update
	if err := updateFunc(cdb.DB.WithContext(ctx)); err != nil {
		return err
	}

	// Invalidate cache
	if len(cacheKeys) > 0 {
		go cdb.Cache.Delete(context.Background(), cacheKeys...)
	}

	return nil
}

// Transaction with cache invalidation
func (cdb *CachedDB) TransactionWithCache(
	ctx context.Context,
	cacheKeys []string,
	fn func(*gorm.DB) error,
) error {
	err := cdb.DB.WithContext(ctx).Transaction(fn)
	if err != nil {
		return err
	}

	// Invalidate cache after successful transaction
	if len(cacheKeys) > 0 {
		go cdb.Cache.Delete(context.Background(), cacheKeys...)
	}

	return nil
}
