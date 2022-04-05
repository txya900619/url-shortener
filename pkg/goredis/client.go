package goredis

import (
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func NewGoRedisClient() redis.UniversalClient {
	var client redis.UniversalClient
	hosts := strings.Split(viper.GetString("REDIS_HOSTS"), ",")
	if len(hosts) <= 1 {
		client = redis.NewClient(&redis.Options{
			Addr:     hosts[0],
			Password: viper.GetString("REDIS_PASSWORD"),
			DB:       0,
		})
	} else {
		client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    hosts,
			Password: viper.GetString("REDIS_PASSWORD"),
		})
	}

	return client
}
