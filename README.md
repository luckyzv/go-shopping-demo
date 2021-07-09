# go-shopping-demo

### Feature
[] logrus。写入日志到文件中，区分不同功能logger
[] viper。加载配置文件
[] auth
[] gin-swagger

### Todo

[] docker
[] rabbitmq死信队列
[] 接入elk
[] 短信邮件推送服务（微服务）
[] grpc 远程调用
[] nodejs api网关
[] consul 服务注册与发现
[] 测试
[] ci/cd


### Deploy

1. cd docker/redis && docker-compose up -d

2. cd ../sentinel && docker-compose up -d

3. cd ../.. && docker-compose up -d
