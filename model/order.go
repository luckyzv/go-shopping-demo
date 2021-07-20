package model

import (
  "gorm.io/gorm"
)

type Order struct {
  gorm.Model

  OrderId string `json:"orderId" gorm:"type:varchar(20);not null;uniqueIndex;"`
  Price float64 `json:"price" gorm:"UNSIGNED NOT NULL;comment:'支付价格'"`
  Version int `json:"version" gorm:"type:tinyint;not null;default 1;unsigned;comment:'版本号'"`
  Status string `json:"status" gorm:"type:enum('created','paid','finished', 'canceled');default:'created'"`

  UserId uint `json:"userId"`
  User User  `json:"user"`
  Products []OrderProduct `json:"orderProducts"`
  OrderHistories []OrderHistory `json:"orderHistories"`
  OrderPayments []OrderPayment `json:"orderPayments"`
}

func OrderIsExistedByOrderId(db *gorm.DB, orderId string) (bool, error) {
  var order Order
  err := db.Select("id").Where("order_id = ?", orderId).First(&order).Error

  if err != nil && err != gorm.ErrRecordNotFound {
    return false, err
  }

  if order.ID > 0 {
    return true, nil
  }

  return false, nil
}

func OrderAddNew(db *gorm.DB, order Order) error {
  err := db.Create(&order).Error
  if err != nil {
    return err
  }

  return nil
}

func OrderGetAll(db *gorm.DB, ids []int) ([]Order, int64)  {
  var orders []Order
  result := db.Find(&orders, ids)

  return orders, result.RowsAffected
}
