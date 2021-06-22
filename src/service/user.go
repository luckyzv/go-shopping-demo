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

func UserRegister(ctx *gin.Context, db *gorm.DB, user model.User)  {
  err := model.CreateUser(db, user)
  if err != nil {
    response.ServerFailedResponse(ctx , constant.ErrorCreateUserFail)
    return
  }
  response.Response(ctx , constant.SUCCESS, nil)
}

func UserLogin(ctx *gin.Context, user model.User) {
  token, err := common.ReleaseToken(user)
  if err != nil {
    response.ServerFailedResponse(ctx , constant.ErrorTokenReleaseFail)
    common.Logger("service", "UserLogin", constant.ErrorTokenReleaseFail, err)
    return
  }

  response.Response(ctx , constant.SUCCESS, dto.UserLoginResponseDto(user, token))
}
