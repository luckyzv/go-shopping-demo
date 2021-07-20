package model

import "gorm.io/gorm"

type OrderPayment struct {
  gorm.Model

  TotalPrice float64  `json:"total_price" gorm:"unsigned not null;comment:'总额'"`
  Discount float64 `json:"discount" gorm:"unsigned not null;comment:'优惠金额'"`
  RealPaid float64 `json:"real_paid" gorm:"unsigned not null;comment:'实付金额'"`
  Status string `json:"status" gorm:"type:enum('pending','finished','canceled');default:'pending'"`

  OrderID uint `json:"orderId"`
}
