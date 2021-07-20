package controller

import (
  "encoding/json"
  "fmt"
  "github.com/gin-gonic/gin"
  "shopping/engine"
  "shopping/model"
  "shopping/response"
  "shopping/response/constant/errorcode"
  "shopping/src/dto"
  "shopping/src/service"
  "strconv"
)

type ProductController struct {}

var productService = &service.ProductService{}

func (c *ProductController) AddNewProduct(ctx *gin.Context)  {
  db := engine.GetMysqlClient()
  var newProductDto dto.NewProductDto

  if err := ctx.ShouldBindJSON(&newProductDto); err != nil {
    response.ClientFailedResponse(ctx, errorcode.ErrorRequiredParamFail)
    return
  }

  existed, _ := model.ProductIsExistedBySkuId(db, newProductDto.SkuId)
  if existed {
    response.ClientFailedResponse(ctx, errorcode.ErrorProductSkuIdDuplicated)
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
  var product model.Product
  db := engine.GetMysqlClient()

  productId := ctx.Param("productId")
  id, _ := strconv.Atoi(productId)

  val := engine.Get(fmt.Sprintf("product:%s:info", productId))
  fmt.Println("val: ", val)
  if val != "" {
    json.Unmarshal([]byte(val), &product)
  } else {
    product = model.ProductGetOne(db, id)
    data, _ := json.Marshal(product)
    engine.Set(fmt.Sprintf("product:%s:info", productId), data, 0)
  }

  response.Response(ctx , errorcode.SUCCESS, product)
}
