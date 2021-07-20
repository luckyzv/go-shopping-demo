package engine

import (
  "fmt"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
  "shopping/config"
  "shopping/model"
  "time"
)

var mysqlClient *gorm.DB

func init()  {
  viperConfig := config.GetMysqlConfig()
  dsn := fmt.Sprintf("%s:%s" + "@tcp(" + "%s:%s" + ")/%s" + "?charset=utf8&parseTime=True&loc=Local",
    viperConfig.UserName, viperConfig.Pass, viperConfig.Host, viperConfig.Port, viperConfig.DbName)

  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil  {
    panic(fmt.Errorf("Connet mysql failed, err: %s \n", err))
  }

  mysqlClient = db

  // 连接池
  sqlDb, err := db.DB()
  if err != nil {
    panic(fmt.Errorf("Connect pool failed, err: %s\n", err))
  }
  sqlDb.SetMaxIdleConns(10)
  sqlDb.SetMaxOpenConns(150)
  sqlDb.SetConnMaxLifetime(time.Hour)

  // 初始化table
  db.AutoMigrate(&model.User{}, &model.Product{}, &model.Order{}, &model.OrderProduct{}, &model.OrderHistory{}, &model.OrderPayment{})

  fmt.Println("初始化数据库成功")
}

func GetMysqlClient() *gorm.DB {
  return mysqlClient
}
