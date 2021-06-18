package controller

import (
  "github.com/gin-gonic/gin"
  "shopping/response"
  "shopping/response/constant"
)

type ProductController struct {

}

func (c *ProductController) AddProduct(ctx *gin.Context)  {
  response.Response(ctx, constant.SUCCESS, nil)
}
