package engine

import (
  "context"
  "github.com/go-redis/redis/v8"
  "shopping/config"
  "time"
)

var redisClient *redis.Client
var ctx = context.Background()

func init()  {
  viperConfig := config.GetRedisConfig()

  // 单机连接
  //redisClient = redis.NewClient(&redis.Options{
  //  Addr: viperConfig.Addr,
  //  Password: viperConfig.Password,
  //  DB: viperConfig.DB,
  //})

  // 哨兵模式连接
  redisClient = redis.NewFailoverClient(&redis.FailoverOptions{
    MasterName: viperConfig.SentinelMasterName,
    SentinelAddrs: viperConfig.SentinelAddr,
    Password: viperConfig.Password,
    DB: 0,
  })
}

func Set(key string, value interface{}, expiration time.Duration)  {
  redisClient.Set(ctx, key, value, expiration)
}

func Get(key string) string {
  return redisClient.Get(ctx, key).Val()
}
