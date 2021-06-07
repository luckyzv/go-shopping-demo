package engine

import (
  "fmt"
  "github.com/streadway/amqp"
  "shopping/config"
)

type RabbitmqClient struct {
  conn *amqp.Connection
  channel *amqp.Channel
}

var rabbitmqClient *RabbitmqClient

func init()  {
  amqpConfig := config.GetAmqpConfig()
  url := fmt.Sprintf("amqp://%s:%s@%s", amqpConfig.UserName, amqpConfig.Password, amqpConfig.Host)
  conn, err := amqp.Dial(url)
  if err != nil {
    panic(fmt.Errorf("Connect rabbitmq failed, err: %s \n", err))
  }
  channel, err := conn.Channel()
  if err != nil {
    panic(fmt.Errorf("Failed to open a channel, err: %s \n", err))
  }
  rabbitmqClient.conn = conn
  rabbitmqClient.channel = channel
}

func CloseRabbitmq() {
  _ = rabbitmqClient.channel.Close()
  _ = rabbitmqClient.conn.Close()
}

func GetRabbitmqConn() *amqp.Connection {
  return rabbitmqClient.conn
}

func GetRabbitmqChannel() *amqp.Channel  {
  return rabbitmqClient.channel
}
