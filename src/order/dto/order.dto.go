package dto

import "shopping/model"

type AddNewOrderDto struct {
  OrderId string `json:"orderId" binding:"required"`
  TotalPrice float64 `json:"totalPrice" binding:"required"`
  OrderProducts []OrderProductsDto `json:"orderProducts" binding:"required"`
}

type OrderProductsDto struct {
  ProductId uint `json:"productId" binding:"required"`
  Number uint `json:"number" binding:"required"`
}

type OrderDto struct {
  ID uint `json:"id"`
  OrderId string `json:"orderId"`
  Price float64 `json:"price"`
  Version int `json:"version"`
  Status string `json:"status"`

  User model.User  `json:"user"`
  OrderProducts []model.OrderProduct `json:"orderProducts"`
  OrderHistories []model.OrderHistory `json:"orderHistories"`
  OrderPayments []model.OrderPayment `json:"orderPayments"`
}
