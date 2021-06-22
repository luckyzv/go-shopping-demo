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

type UserController struct {

}

// UserRegister @用户注册
// @Tags Users
// @Description 用户注册
// @Accept json
// @Produce json
// @Param phone body string  true "注册手机号"
// @Param password body string  true "用户密码"
// @Success 200 {string} string "ok"
// @Router /users/register [post]
func (c *UserController) UserRegister(ctx *gin.Context) {
  db := engine.GetMysqlClient()

  var user dto.UserLoginDto
  if err := ctx.ShouldBindJSON(&user); err != nil {
    response.ClientFailedResponse(ctx, constant.ErrorRequiredParamFail)
    return
  }

  // 手机号已被注册
  existed, _ := model.ExistUserByPhone(db, user.Phone)
  if existed {
    response.ClientFailedResponse(ctx, constant.ErrorUserExisted)
    return
  }
  // 加密password
  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
  if err != nil {
    response.ServerFailedResponse(ctx, constant.ErrorHashedPasswordFail)
    return
  }
  // 生成用户
  randomNameByteLength := 8
  service.UserRegister(ctx, db, model.User{
    Name: util.GetRandomString(randomNameByteLength),
    PassWord: string(hashedPassword),
    Phone: user.Phone,
  })
}

// UserLogin 用户登录
func (c *UserController) UserLogin(ctx *gin.Context) {
  db := engine.GetMysqlClient()

  var user dto.UserLoginDto
  if err := ctx.ShouldBindJSON(&user); err != nil {
    response.ClientFailedResponse(ctx, constant.ErrorRequiredParamFail)
    return
  }

  // 该手机号不存在
  isExisted, dbUser := model.GetUserByPhone(db, user.Phone)
  if !isExisted {
    response.ClientFailedResponse(ctx, constant.ErrorUserNonExisted)
    return
  }
  // 比对密码失败
  err := bcrypt.CompareHashAndPassword([]byte(dbUser.PassWord), []byte(user.Password))
  if err != nil {
    response.ClientFailedResponse(ctx, constant.ErrorPasswordCheckFail)
    return
  }
  service.UserLogin(ctx, *dbUser)
}

func (c *UserController) UserInfo(ctx *gin.Context)  {
  db := engine.GetMysqlClient()
  userId, _ := ctx.Get("userId")
  user, err := model.GetUserById(db, userId.(uint))
  if err != nil {
    response.ClientFailedResponse(ctx, constant.ErrorUserNonExisted)
    return
  }

  response.Response(ctx, constant.SUCCESS, dto.ConvertModelUserToDto(*user))
}
