package dto

import "time"

type ProductDto struct {
  Id uint `json:"id"`
  CreatedAt time.Time `json:"createdAt"`
  DeletedAt time.Time  `json:"deletedAt"`
  NewProductDto
  Status string `json:"status"`
}

type NewProductDto struct {
  SkuId string `json:"skuId" binding:"required"`
  SkuName string `json:"skuName" binding:"required"`
  Price float64 `json:"price" binding:"required"`
  PromotionPrice float64 `json:"promotionPrice"`
  Stock int `json:"stock" binding:"required"`
}

func ProductDtoTo()  {

}
