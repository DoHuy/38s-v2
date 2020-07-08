package clients

import (
	"github.com/go-redis/redis"
)

func NewRedis(options *redis.Options) (*redis.Client, error) {
	client := redis.NewClient(options)
	if _, err := client.Ping().Result(); err != nil {
		return nil, err
	}

	return client, nil
}
