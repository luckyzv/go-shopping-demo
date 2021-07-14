package service

import (
  "github.com/gin-gonic/gin"
  "gorm.io/gorm"
  "shopping/model"
  "shopping/response"
  "shopping/response/constant"
  "shopping/src/dto"
)

type ProductService struct {
  dto.ProductDto

  PageNum int
  PageSize int
}

func (productService *ProductService) AddNewProduct(ctx *gin.Context, db *gorm.DB, product model.Product)  {
  err := model.ProductAddNew(db, product)
  if err != nil {
    response.ServerFailedResponse(ctx , constant.ErrorProductCreateProductFail)
    return
  }
  response.Response(ctx , constant.SUCCESS, nil)
}

func (productService *ProductService) GetMaps() map[string]interface{}  {
  maps := make(map[string]interface{})
  maps["deleted_on"] = 0

  if productService.SkuId != "" {
    maps["skuId"] = productService.SkuId
  }
  if productService.SkuName != "" {
    maps["skuName"] = productService.SkuName
  }
  if productService.Status != "" {
    maps["status"] = productService.Status
  }

  return maps
}