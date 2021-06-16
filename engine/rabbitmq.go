package engine

import (
  "fmt"
  "github.com/streadway/amqp"
  "shopping/config"
)

var rabbitmqConn *amqp.Connection

func init()  {
  amqpConfig := config.GetAmqpConfig()
  url := fmt.Sprintf("amqp://%s:%s@%s/", amqpConfig.UserName, amqpConfig.Password, amqpConfig.Host)

  conn, err := amqp.Dial(url)
  if err != nil {
    panic(fmt.Errorf("Connect rabbitmq failed, err: %s \n", err))
  }

  rabbitmqConn = conn
}

func CloseRabbitmq() {
  _ = rabbitmqConn.Close()
}

func GetRabbitmqConn() *amqp.Connection {
  return rabbitmqConn
}
