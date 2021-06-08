package config

import (
  "fmt"
  "github.com/spf13/viper"
  "os"
)

type Config struct {
  Env string `json:"env"`
  Version string `json:"version"`
  Mysql MysqlConfig `mapstructure:"mysql"`
  Redis RedisConfig  `mapstructure:"redis"`
  Server ServerConfig `mapstructure:"server"`
  Amqp AmqpConfig `mapstructure:"amqp"`
}

type MysqlConfig struct {
  UserName string `json:"userName"`
  Pass string `json:"pass"`
  Host string `json:"host"`
  Port string `json:"port"`
  DbName string `json:"dbName"`
}

type RedisConfig struct {
  Addr string `json:"addr"`
  Password string `json:"password"`
  DB int `json:"db"`
  SentinelMasterName string `json:"sentinelMasterName"`
  SentinelAddr []string `json:"sentinel-addr"`
}

type AmqpConfig struct {
  UserName string  `json:"userName"`
  Password string `json:"password"`
  Host string  `json:"host"`
}

type ServerConfig struct {
  Port string `json:"port"`
}

var viperConfig Config

func init()  {
  workDir, _ := os.Getwd()
  EnvName := os.Getenv("ENV")
  if EnvName == "" {
    EnvName = "dev"
  }
  viper.SetConfigName("conf." + EnvName)
  viper.SetConfigType("yaml")
  viper.AddConfigPath(workDir + "/config")

  err := viper.ReadInConfig()
  if err != nil {
    panic(fmt.Errorf("Fatal error config file: %s \n", err))
  }

  if err = viper.Unmarshal(&viperConfig); err != nil {
    panic(fmt.Errorf("Unmarshal conf failed, err: %s \n", err))
  }

  // 监控文件变化，不需要重新启动服务器
  //viper.WatchConfig()
  //viper.OnConfigChange(func(in fsnotify.Event) {
  // fmt.Println("配置文件修改...")
  // if err := viper.Unmarshal(&viperConfig); err != nil {
  //   panic(fmt.Errorf("Unmarshal conf failed, err: %s \n", err))
  // }
  //})
}

func GetServerConfig() ServerConfig  {
  return viperConfig.Server
}

func GetMysqlConfig() MysqlConfig  {
  return viperConfig.Mysql
}

func GetRedisConfig() RedisConfig  {
  return viperConfig.Redis
}

func GetAmqpConfig() AmqpConfig {
  return viperConfig.Amqp
}
