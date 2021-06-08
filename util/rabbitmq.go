package util

import (
  "fmt"
  "github.com/streadway/amqp"
  "log"
  "shopping/engine"
)

func PublishMessage(queueName string, msg string)  {
  channel, err := engine.GetRabbitmqConn().Channel()
  defer channel.Close()
  FailOnError(err, "Failed to open a channel")

  // 持久化队列
  _, err = channel.QueueDeclare(queueName, true, false, false, false, nil)
  FailOnError(err, "Failed to declare a queue")

  err = channel.Publish("", queueName, false, false, amqp.Publishing{
    DeliveryMode: amqp.Persistent, // 持久化该消息
    ContentType:  "text/plain",
    Body:         []byte(msg),
  })
  FailOnError(err, fmt.Sprintf("Failed to publish a message, message: %s", msg))
}

func ConsumeMessage(queueName string, userId string)  {
  channel, err := engine.GetRabbitmqConn().Channel()
  FailOnError(err, "Failed to open a channel")

  _, err = channel.QueueDeclare(queueName, true, false, false, false, nil)
  FailOnError(err, "Failed to declare a queue")

  messages, err := channel.Consume(queueName, "", false, false, false, false, nil)
  FailOnError(err, "Failed to consume a message")

  forever := make(chan bool)
  go func() {
    for message := range messages {
      message.Ack(false) // 手动确认该消息已被消费，false表示只针对当前消息
      log.Printf("接收端: %s ======= [x] %s", userId, message.Body)
      // 业务代码
    }
  }()
  <-forever
}

func FailOnError(err error, msg string)  {
  if err != nil {
    panic(fmt.Errorf("%s, err: %s \n", msg, err))
  }
}
