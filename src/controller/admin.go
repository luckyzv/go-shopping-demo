package controller

import (
  "github.com/gin-gonic/gin"
  "shopping/response"
  "shopping/response/constant"
)

type AdminController struct {

}

func (c *AdminController) Login(ctx *gin.Context)  {
  response.Response(ctx, constant.SUCCESS, nil)
}

func (c *AdminController) CreateProduct(ctx *gin.Context)  {
  response.Response(ctx, constant.SUCCESS, nil)
}
