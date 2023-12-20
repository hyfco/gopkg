package bootstrap

import (
	"github.com/hyfco/gopkg/config"
	"time"
)

func NewRedisClient(conf *config.Data) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Network:      conf.Redis.Network,
		DialTimeout:  time.Duration(5 * time.Second),
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
	})
}
