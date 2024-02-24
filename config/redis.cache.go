package config

type RedisConfig struct {
	Hostname string
	Port     string
}

func NewRedisConfig(host string, port string) *RedisConfig {
	return &RedisConfig{
		Hostname: host,
		Port:     port,
	}
}
