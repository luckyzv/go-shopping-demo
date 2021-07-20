package order

import (
  "fmt"
  "github.com/streadway/amqp"
  "log"
  "shopping/config"
  "shopping/engine"
  "strings"
)

func init()  {
  fmt.Println("开始监听死信队列")
  go ConsumerDlx()
}

func PublishOrderMessage(msg string)  {
  keyConfig := config.GetAmqpKeyConfig()
  channel, err := engine.GetRabbitmqConn().Channel()
  defer channel.Close()
  FailOnError(err, "Failed to open a channel")

  // 持久化队列
  err = channel.ExchangeDeclare(keyConfig.OrderExchange, amqp.ExchangeDirect, true, false, false, false, nil)
  FailOnError(err, "Failed to declare a exchange")

  err = channel.Publish(keyConfig.OrderExchange, keyConfig.OrderRoutingKey, false, false, amqp.Publishing{
    DeliveryMode: amqp.Persistent, // 持久化该消息
    ContentType:  "text/plain",
    Body:         []byte(msg),
  })
  FailOnError(err, fmt.Sprintf("Failed to publish a message, message: %s", msg))
}

func consumeOrderMessage()  {
  keyConfig := config.GetAmqpKeyConfig()
  fmt.Println("初始化死信队列")

  channel, err := engine.GetRabbitmqConn().Channel()
  defer channel.Close()
  FailOnError(err, "Failed to open a channel")

  err = channel.ExchangeDeclare(keyConfig.OrderExchange, amqp.ExchangeDirect, true, false, false, false, nil)
  FailOnError(err, "Failed to declare a exchange")

  argsQue := make(map[string]interface{})
  argsQue["x-dead-letter-exchange"] = keyConfig.OrderDlxExchange
  argsQue["x-dead-letter-routing-key"] = keyConfig.OrderDlxRoutingKey
  argsQue["x-message-ttl"] = 1000 * 6 // 毫秒

  queue, err := channel.QueueDeclare(keyConfig.OrderQueue, true, false, false, false, argsQue)
  FailOnError(err, "Failed to declare a queue")

  err = channel.QueueBind(queue.Name, "orders", keyConfig.OrderExchange, false, nil)
  FailOnError(err, "Failed to bind a queue")

  //messages, err := channel.Consume(queue.Name, "", false, false, false, false, nil)
  //FailOnError(err, "Failed to consume a message")

  //forever := make(chan bool)
  //go func() {
  // for message := range messages {
  //   message.Ack(false) // 手动确认该消息已被消费，false表示只针对当前消息
  //   log.Printf("接收端: ======= %s", message.Body)
  //   // 业务代码
  // }
  //}()
  //<-forever
}

func ConsumerDlx()  {
  consumeOrderMessage()

  keyConfig := config.GetAmqpKeyConfig()
  channel, err := engine.GetRabbitmqConn().Channel()
  defer channel.Close()
  FailOnError(err, "Failed to open a channel")

  err = channel.ExchangeDeclare(keyConfig.OrderDlxExchange, amqp.ExchangeDirect, true, false, false, false, nil)
  FailOnError(err, "Failed to declare a exchange")

  queue, err := channel.QueueDeclare(keyConfig.OrderDlxQueue, true, false, false, false, nil)
  FailOnError(err, "Failed to declare a queue")

  err = channel.QueueBind(queue.Name, keyConfig.OrderDlxRoutingKey, keyConfig.OrderDlxExchange, false, nil)
  FailOnError(err, "Failed to bind a queue")

  messages, err := channel.Consume(queue.Name, "", false, false, false, false, nil)
  FailOnError(err, "Failed to consume a message")

  forever := make(chan bool)
  go func() {
    for message := range messages {
      // 业务代码 库存增加，设置订单状态为取消
      str := string(message.Body)
      index := strings.Index(str, "[")
      orderId := str[index + 1 : len(str) - 1]

      UpdateOrderStatus(orderId)

      message.Ack(false) // 手动确认该消息已被消费，false表示只针对当前消息
      log.Printf("死信队列接收端: ======= [x] %s", message.Body)
    }
  }()
  <-forever

}

func FailOnError(err error, msg string)  {
  if err != nil {
    panic(fmt.Errorf("%s, err: %s \n", msg, err))
  }
}
