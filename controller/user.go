package controller

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "shopping/response"
  "shopping/response/constant"
  "shopping/dto"
  "shopping/engine"
  "shopping/model"
  "shopping/service"
  "shopping/util"
)

// UserRegister 用户注册
func UserRegister(c *gin.Context) {
  db := engine.GetMysqlClient()

  var user dto.UserLoginDto
  if err := c.ShouldBind(&user); err != nil {
    response.FailedResponse(c, constant.Error_RequiredParamFail)
    return
  }
  existed, _ := model.ExistUserByPhone(db, user.Phone)
  if existed {
    response.FailedResponse(c, constant.Error_UserExistedFail)
    return
  }

  randomNameByteLength := 8
  service.UserRegister(c, db, model.User{
    Name: util.GetRandomString(randomNameByteLength),
    PassWord: user.Password,
    Phone: user.Phone,
  })
}

// UserLogin 用户登录
func UserLogin(c *gin.Context) {
  response.Response(c, http.StatusOK, constant.SUCCESS, nil)
}
