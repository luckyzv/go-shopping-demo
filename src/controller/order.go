package controller

import (
  "github.com/gin-gonic/gin"
  "shopping/response"
  "shopping/response/constant"
)

type OrderController struct {

}

func (c *OrderController) CreateOrder(ctx *gin.Context)  {
  response.Response(ctx, constant.SUCCESS, nil)
}
