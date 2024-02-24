package services

import (
	"context"

	"github.com/spoonbuoy/lcdb/interfaces/icache"
)

type redis struct {
	cacherepo icache.ICacheRepo
}

func newRedis(cache icache.ICacheRepo) *redis {
	return &redis{
		cacherepo: cache,
	}
}

func (r *redis) GetCacheKey(ctx context.Context, cacheKey string) (string, error) {
	_, _ = r.cacherepo.GetCacheKey(ctx, cacheKey)
	return "", nil
}
