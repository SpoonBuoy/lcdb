package repository

import (
	"context"

	"github.com/go-redis/redis"
)

type redisCacheRepo struct {
	client *redis.Client
}

func newCacher(client *redis.Client) *redisCacheRepo {
	return &redisCacheRepo{
		client: client,
	}
}

func (rcr *redisCacheRepo) GetCacheKey(ctx context.Context, cacheKey string) (string, error) {
	return rcr.client.Get(cacheKey).Val(), nil
}
