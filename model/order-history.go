package model

import "gorm.io/gorm"

type OrderHistory struct {
  gorm.Model

  Status string  `json:"status" gorm:"type:enum('created','paid','finished','canceled');default:'created';uniqueIndex:orderId_status""`
  OrderID uint `json:"orderId" gorm:"uniqueIndex:orderId_status"`
}
