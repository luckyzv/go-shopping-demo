package model

import "gorm.io/gorm"

type Product struct {
  gorm.Model

  SkuId string `json:"skuId" gorm:"size:11;not null;uniqueIndex"`
  SkuName string `json:"skuName" gorm:"type:varchar(255);not null"`
  Price string `json:"price" gorm:"size:11;not null"`
}
