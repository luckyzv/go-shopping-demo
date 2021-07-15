package controller

import (
  "github.com/gin-gonic/gin"
  "shopping/engine"
  "shopping/model"
  "shopping/response"
  "shopping/response/constant"
  "shopping/src/dto"
  "shopping/src/service"
)

type ProductController struct {}

var productService = &service.ProductService{}

func (c *ProductController) AddProduct(ctx *gin.Context)  {
  db := engine.GetMysqlClient()
  var newProductDto dto.NewProductDto

  if err := ctx.ShouldBindJSON(&newProductDto); err != nil {
    response.ClientFailedResponse(ctx, constant.ErrorRequiredParamFail)
    return
  }

  existed, _ := model.ProductIsExistedBySkuId(db, newProductDto.SkuId)
  if existed {
    response.ClientFailedResponse(ctx, constant.ErrorProductSkuIdDuplicated)
    return
  }

  product := model.Product{
    SkuId: newProductDto.SkuId,
    SkuName: newProductDto.SkuName,
    Price: newProductDto.Price,
    Stock: newProductDto.Stock,
  }
  if newProductDto.PromotionPrice > 0 {
    product.PromotionPrice = newProductDto.PromotionPrice
  }
  productService.AddNewProduct(ctx, db, product)
}

func (c *ProductController) GetProduct(ctx *gin.Context)  {
  response.Response(ctx, constant.SUCCESS, nil)
}
