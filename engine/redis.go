package engine

import (
  "github.com/go-redis/redis/v8"
  "shopping/config"
  "time"
)

var redisClient *redis.Client

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

func GetRedisClient() *redis.Client {
  return redisClient
}

func Set(key string, value interface{}, expiration time.Duration)  {
  redisClient.Set(redisClient.Context(), key, value, expiration)
}

func Get(key string) string {
  return redisClient.Get(redisClient.Context(), key).Val()
}
