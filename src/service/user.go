package service

import (
  "github.com/gin-gonic/gin"
  "gorm.io/gorm"
  "shopping/common"
  "shopping/middleware"
  "shopping/model"
  "shopping/response"
  "shopping/response/constant"
  "shopping/src/dto"
)

func UserRegister(c *gin.Context, db *gorm.DB, user model.User)  {
  err := model.CreateUser(db, user)
  if err != nil {
    response.ServerFailedResponse(c, constant.ErrorCreateUserFail)
    return
  }
  response.Response(c, constant.SUCCESS, nil)
}

func UserLogin(c *gin.Context, user model.User) {
  token, err := common.ReleaseToken(user)
  if err != nil {
    response.ServerFailedResponse(c, constant.ErrorTokenReleaseFail)
    middleware.Logger().Error("服务器内部异常：", err)
    return
  }

  response.Response(c, constant.SUCCESS, dto.UserInfo(user, token))
}
