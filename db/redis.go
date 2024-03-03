package db

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis"
	"github.com/spoonbuoy/lcdb/config"
)

var (
	client *redis.Client
	doOnce sync.Once
)

func initRedis() {
	// todo
	conf := config.RedisConfig{}

	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf(`%s, %s`, conf.Hostname, conf.Port),
		Password: "", // no password set
		DB:       0,  // use default c
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

}

func GetRedisClient() *redis.Client {
	if client == nil {
		doOnce.Do(initRedis)
	}
	return client
}

type Cache struct {
	Config config.RedisConfig
}

func (c *Cache) Connect() (Store, error) {
	return c, nil
}

func (c *Cache) Close() error {
	return nil
}
func (c *Cache) GetClient() error {
	return nil
}
