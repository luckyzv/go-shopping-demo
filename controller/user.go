package controller

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "shopping/common"
  "shopping/common/constant"
  "shopping/dto"
  "shopping/engine"
  "shopping/model"
)

// UserRegister 用户注册
func UserRegister(c *gin.Context) {
  db := engine.GetMysqlClient()

  var user dto.UserLoginDto
  if err := c.ShouldBindJSON(&user); err != nil {
    common.FailedParamResponse(c, constant.ErrorRequiredParamFail)
  }
  existed, _ := model.ExistUserByEmail(db, user.Email)
  if existed {
    common.FailedParamResponse(c, constant.ErrorUserExistedFail)
  }

  //service.UserRegister(c, &user)
}

// UserLogin 用户登录
func UserLogin(c *gin.Context) {
  //logger := middleware.GetLogger()
  //logger.Debug("oh 内容%d ")
  common.Response(c, http.StatusOK, constant.SUCCESS, nil)
}
