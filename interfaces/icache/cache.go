package icache

import (
	"context"
	"time"
)

type ICacheService interface {
	SetCacheKey(ctx context.Context, cacheKey, cacheValue string, expiration time.Duration) error
	GetCacheKey(ctx context.Context, cacheKey string) (string, error)
	DeleteCacheKey(ctx context.Context, cacheKey string) error
}

type ICacheRepo interface {
	SetCacheKey(ctx context.Context, cacheKey, cacheValue string, expiration time.Duration) error
	GetCacheKey(ctx context.Context, cacheKey string) (string, error)
	DeleteCacheKey(ctx context.Context, cacheKey string) error
}
