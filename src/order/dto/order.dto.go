package dto

type AddNewOrderDto struct {
  OrderId string `json:"orderId" binding:"required"`
  TotalPrice float64 `json:"totalPrice" binding:"required"`
  OrderProducts []OrderProductsDto `json:"orderProducts" binding:"required"`
}

type OrderProductsDto struct {
  ProductId uint `json:"productId" binding:"required"`
  Number uint `json:"number" binding:"required"`
}
