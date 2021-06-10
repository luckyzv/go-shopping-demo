package service

import (
  "github.com/gin-gonic/gin"
  "gorm.io/gorm"
  "net/http"
  "shopping/common"
  "shopping/common/constant"
  "shopping/dto"
)

func UserRegister(c *gin.Context, db *gorm.DB, user *dto.UserDto)  {
  //model.CreateUser(db, user)
  common.Response(c, http.StatusOK, constant.SUCCESS, nil)
}
