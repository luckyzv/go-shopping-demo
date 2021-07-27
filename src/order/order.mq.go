package order

import (
  "fmt"
  "github.com/streadway/amqp"
  "gorm.io/gorm"
  "log"
  "shopping/config"
  "shopping/engine"
  "shopping/model"
  "strings"
)

func init()  {
  fmt.Println("开始监听死信队列")
  go ConsumerDlx(AtomicDecrStock)
}

type Callback func(orderId string)

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

func declareNormalConsumerOrderMessage()  {
  keyConfig := config.GetAmqpKeyConfig()

  channel, err := engine.GetRabbitmqConn().Channel()
  defer channel.Close()
  FailOnError(err, "Failed to open a channel")

  // 最多只能允许接收三个未确认的
  err = channel.Qos(3, 0, false)
  FailOnError(err, "Failed to setting qos")

  err = channel.ExchangeDeclare(keyConfig.OrderExchange, amqp.ExchangeDirect, true, false, false, false, nil)
  FailOnError(err, "Failed to declare a exchange")

  argsQue := make(map[string]interface{})
  argsQue["x-dead-letter-exchange"] = keyConfig.OrderDlxExchange
  argsQue["x-dead-letter-routing-key"] = keyConfig.OrderDlxRoutingKey
  argsQue["x-message-ttl"] = 1000 * 30 // 单位：毫秒， 30s

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

func ConsumerDlx(callback Callback)  {
  declareNormalConsumerOrderMessage()

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

      callback(orderId)

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

func AtomicDecrStock(orderId string)  {
  db := engine.GetMysqlClient()

  // 事务中进行
  db.Transaction(func(tx *gorm.DB) error {
    var order model.Order
    tx.Model(&model.Order{}).Where("orders.order_id = ?", orderId).Find(&order)

    if order.Status == "created" {
      // 修改为取消状态
      tx.Model(&order).Update("status", "canceled")

      // 加库存
      var orderProducts []model.OrderProduct
      tx.Model(&order).Association("OrderProducts").Find(&orderProducts)
      for _, orderProduct := range orderProducts {
        tx.Model(&model.Product{}).Where("id = ?", orderProduct.ProductID).Update("stock", gorm.Expr("stock + ?", orderProduct.Num))
      }
    }

    return nil
  })
}
