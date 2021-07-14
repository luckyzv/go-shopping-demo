package controller

import (
  "github.com/gin-gonic/gin"
  "shopping/common"
  "shopping/engine"
  "shopping/response"
  "shopping/response/constant"
  "shopping/src/dto"
  "shopping/src/service"
)

type AdminController struct {

}

func (c *AdminController) Login(ctx *gin.Context)  {
  response.Response(ctx, constant.SUCCESS, nil)
}

func (c *AdminController) GetALlUsers(ctx *gin.Context) {
  db := engine.GetMysqlClient()

  var adminDto dto.AdminDto

  err := ctx.ShouldBindJSON(&adminDto)
  if err != nil {
    response.ClientFailedResponse(ctx, constant.ErrorRequiredParamFail)
    return
  }

  allUsers, err := service.UserGetAll(db, adminDto.PageNum, adminDto.PageSize, adminDto.Status)
  if err != nil {
    response.ServerFailedResponse(ctx, constant.ErrorUserFindFail)
    common.Logger("AdminController", "GetAllUsers", constant.ErrorUserFindFail, err)
    return
  }

  response.Response(ctx, constant.SUCCESS, allUsers)
}

func (c *AdminController) CreateProduct(ctx *gin.Context)  {
  response.Response(ctx, constant.SUCCESS, nil)
}
