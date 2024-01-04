package bootstrap

import (
	"github.com/hyfco/gopkg/conf"
	"github.com/redis/go-redis/v9"
	"time"
)

func NewRedisClient(conf *conf.Bootstrap) *redis.Client {
	redisConf := conf.GetData().GetRedis()
	return redis.NewClient(&redis.Options{
		Addr:         redisConf.GetAddr(),
		Network:      redisConf.Network,
		DialTimeout:  time.Duration(5 * time.Second),
		ReadTimeout:  redisConf.GetReadTimeout().AsDuration(),
		WriteTimeout: redisConf.GetWriteTimeout().AsDuration(),
	})
}
