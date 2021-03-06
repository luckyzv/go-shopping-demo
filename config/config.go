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
  Logger LoggerConfig `mapstructure:"logger"`
  Jwt JwtConfig `mapstructure:"jwt"`
  Es EsConfig `mapstructure:"elasticsearch"`
  AmqpKey AmqpKeyConfig `mapstructure:"amqpkey"`
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

type AmqpKeyConfig struct {
  OrderDlxExchange string `json:"OrderDlxExchange"`
  OrderDlxQueue string `json:"OrderDlxQueue"`
  OrderDlxRoutingKey string `json:"OrderDlxRoutingKey"`
  OrderExchange string `json:"OrderExchange"`
  OrderQueue string `json:"OrderQueue"`
  OrderRoutingKey string `json:"orderRoutingKey"`
}

type EsConfig struct {
  Host string `json:"host"`
  Port string `json:"port"`
}

type LoggerConfig struct {
  AppFilePath string `json:"AppFilePath"`
  ApiFilePath string `json:"ApiFilePath"`
}

type ServerConfig struct {
  Port string `json:"port"`
}

type JwtConfig struct {
  JwtKey string `json:"jwtKey"`
  Issuer string `json:"issuer"`
  Subject string `json:"subject"`
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

  // ???????????????????????????????????????????????????
  //viper.WatchConfig()
  //viper.OnConfigChange(func(in fsnotify.Event) {
  // fmt.Println("??????????????????...")
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

func GetElasticSearchConfig() EsConfig {
  return viperConfig.Es
}

func GetLoggerConfig() LoggerConfig {
  return viperConfig.Logger
}

func GetJwtConfig() JwtConfig {
  return viperConfig.Jwt
}

func GetAmqpKeyConfig() AmqpKeyConfig {
  return viperConfig.AmqpKey
}
