package controller

import (
  "github.com/gin-gonic/gin"
  "golang.org/x/crypto/bcrypt"
  "shopping/common"
  "shopping/engine"
  "shopping/model"
  "shopping/response"
  "shopping/response/constant"
  "shopping/src/dto"
  "shopping/src/service"
  "shopping/util"
)

type UserController struct {}

var userService = &service.UserService{}

// UserRegister @用户注册
// @Tags Users
// @Description 用户注册
// @Accept json
// @Produce json
// @Param phone body string  true "注册手机号"
// @Param password body string  true "用户密码"
// @Success 200 {object} response.SuccessResBody "ok"
// @Failure 500 {object} response.FailedResBody "internal server error"
// @Router /users/register [post]
func (c *UserController) UserRegister(ctx *gin.Context) {
  db := engine.GetMysqlClient()

  var user dto.UserLoginDto
  if err := ctx.ShouldBindJSON(&user); err != nil {
    response.ClientFailedResponse(ctx, constant.ErrorRequiredParamFail)
    return
  }

  // 手机号已被注册
  existed, _ := model.UserIsExistedByPhone(db, user.Phone)
  if existed {
    response.ClientFailedResponse(ctx, constant.ErrorUserExisted)
    return
  }
  // 加密password
  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
  if err != nil {
    response.ServerFailedResponse(ctx, constant.ErrorUserHashedPasswordFail)
    return
  }
  // 生成用户
  randomNameByteLength := 8
  userService.UserRegister(ctx, db, model.User{
    Name: util.GetRandomString(randomNameByteLength),
    PassWord: string(hashedPassword),
    Phone: user.Phone,
  })
}

// UserLogin @用户登录
// @Tags Users
// @Description 用户登录
// @Accept json
// @Produce json
// @Param phone body string true "手机号"
// @Param password body string true "密码"
// @Success 200 {object} response.SuccessResBody "ok"
// @Failure 500 {object} response.FailedResBody "internal server error"
// @Router /users/login [post]
func (c *UserController) UserLogin(ctx *gin.Context) {
  db := engine.GetMysqlClient()

  var user dto.UserLoginDto
  if err := ctx.ShouldBindJSON(&user); err != nil {
    response.ClientFailedResponse(ctx, constant.ErrorRequiredParamFail)
    return
  }

  // 该手机号不存在
  isExisted, dbUser := model.UserGetOneByPhone(db, user.Phone)
  if !isExisted {
    response.ClientFailedResponse(ctx, constant.ErrorUserNonExisted)
    return
  }
  // 比对密码失败
  err := bcrypt.CompareHashAndPassword([]byte(dbUser.PassWord), []byte(user.Password))
  if err != nil {
    response.ClientFailedResponse(ctx, constant.ErrorUserPasswordCheckFail)
    return
  }
  userService.UserLogin(ctx, *dbUser)
}

func (c *UserController) UserInfo(ctx *gin.Context)  {
  db := engine.GetMysqlClient()
  userId, _ := ctx.Get("userId")
  user, err := model.UserGetOneById(db, userId.(uint))
  if err != nil {
    response.ClientFailedResponse(ctx, constant.ErrorUserNonExisted)
    return
  }

  response.Response(ctx, constant.SUCCESS, dto.ConvertModelUserToDto(*user))
}

func (c *UserController) GetUsers(ctx *gin.Context)  {
  db := engine.GetMysqlClient()
  var getAllUsersDto dto.GetAllUsersDto
  err := ctx.ShouldBindJSON(&getAllUsersDto)
  if err != nil {
    response.ClientFailedResponse(ctx, constant.ErrorRequiredParamFail)
    return
  }

  users, err := userService.GetAllUsers(db, getAllUsersDto)
  if err != nil {
    response.ServerFailedResponse(ctx , constant.ErrorUserFindFail)
    common.Logger("UserService", "GetAllUsers", constant.ErrorUserFindFail, err)
    return
  }

  response.Response(ctx, constant.SUCCESS, users)
}
