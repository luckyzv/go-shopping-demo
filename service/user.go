package service

import (
  "github.com/gin-gonic/gin"
  "gorm.io/gorm"
  "net/http"
  "shopping/response"
  "shopping/response/constant"
  "shopping/model"
)

func UserRegister(c *gin.Context, db *gorm.DB, user model.User)  {
  err := model.CreateUser(db, user)
  if err != nil {
    response.FailedResponse(c, constant.Error_MysqlCreateUserError)
    return
  }
  response.Response(c, http.StatusOK, constant.SUCCESS, nil)
}
