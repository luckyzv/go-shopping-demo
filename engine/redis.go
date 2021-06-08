package engine

import (
  "github.com/go-redis/redis/v8"
  "shopping/config"
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
  })
}

func GetRedisClient() *redis.Client {
  return redisClient
}
