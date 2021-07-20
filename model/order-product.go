package model

import (
  "gorm.io/gorm"
)

type OrderProduct struct {
  gorm.Model

  Price float64 `json:"price" gorm:"comment:'购物时的价格'"`
  Num uint `json:"num" gorm:"type:smallint unsigned not null;comment:'该产品购买数量'"`

  ProductID uint `json:"productId"`
  OrderID uint `json:"orderId"`
}
