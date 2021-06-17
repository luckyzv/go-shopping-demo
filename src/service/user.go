package service

import (
  "github.com/gin-gonic/gin"
  "gorm.io/gorm"
  "shopping/common"
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
    common.Logger("service", "UserLogin", constant.ErrorTokenReleaseFail, err)
    return
  }

  response.Response(c, constant.SUCCESS, dto.UserInfo(user, token))
}
