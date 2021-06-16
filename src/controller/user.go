package controller

import (
  "github.com/gin-gonic/gin"
  "golang.org/x/crypto/bcrypt"
  "shopping/engine"
  "shopping/model"
  "shopping/response"
  "shopping/response/constant"
  "shopping/src/dto"
  "shopping/src/service"
  "shopping/util"
)

// UserRegister 用户注册
func UserRegister(c *gin.Context) {
  db := engine.GetMysqlClient()

  var user dto.UserLoginDto
  if err := c.ShouldBindJSON(&user); err != nil {
    response.ClientFailedResponse(c, constant.ErrorRequiredParamFail)
    return
  }

  // 手机号已被注册
  existed, _ := model.ExistUserByPhone(db, user.Phone)
  if existed {
    response.ClientFailedResponse(c, constant.ErrorUserExisted)
    return
  }
  // 加密password
  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
  if err != nil {
    response.ServerFailedResponse(c, constant.ErrorHashedPasswordFail)
    return
  }
  // 生成用户
  randomNameByteLength := 8
  service.UserRegister(c, db, model.User{
    Name: util.GetRandomString(randomNameByteLength),
    PassWord: string(hashedPassword),
    Phone: user.Phone,
  })
}

// UserLogin 用户登录
func UserLogin(c *gin.Context) {
  db := engine.GetMysqlClient()

  var user dto.UserLoginDto
  if err := c.ShouldBindJSON(&user); err != nil {
    response.ClientFailedResponse(c, constant.ErrorRequiredParamFail)
    return
  }

  // 该手机号不存在
  isExisted, dbUser := model.GetUserByPhone(db, user.Phone)
  if !isExisted {
    response.ClientFailedResponse(c, constant.ErrorUserNonExisted)
    return
  }
  // 比对密码失败
  err := bcrypt.CompareHashAndPassword([]byte(dbUser.PassWord), []byte(user.Password))
  if err != nil {
    response.ClientFailedResponse(c, constant.ErrorPasswordCheckFail)
    return
  }
  service.UserLogin(c, *dbUser)
}
