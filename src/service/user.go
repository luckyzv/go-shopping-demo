package service

import (
  "github.com/gin-gonic/gin"
  "gorm.io/gorm"
  "shopping/common"
  "shopping/model"
  "shopping/response"
  "shopping/response/constant/errorcode"
  "shopping/src/dto"
)

type UserService struct {
  dto.UserDto

  PageNum int
  PageSize int
}

func (userService *UserService) UserRegister(ctx *gin.Context, db *gorm.DB, user model.User)  {
  err := model.UserAddNew(db, user)
  if err != nil {
    response.ServerFailedResponse(ctx , errorcode.ErrorUserCreateUserFail)
    return
  }
  response.Response(ctx , errorcode.SUCCESS, nil)
}

func (userService *UserService) UserLogin(ctx *gin.Context, user model.User) {
  token, err := common.ReleaseToken(user)
  if err != nil {
    response.ServerFailedResponse(ctx , errorcode.ErrorUserTokenReleaseFail)
    common.Logger("UserService", "UserLogin", errorcode.ErrorUserTokenReleaseFail, err, user)
    return
  }

  response.Response(ctx, errorcode.SUCCESS, dto.UserLoginResponseDto(user, token))
}

func (userService *UserService) GetAllUsers(db *gorm.DB, usersDto dto.GetAllUsersDto) ([]model.User, error) {
  allUsers, err := UserGetAll(db, usersDto.PageNum, usersDto.PageSize, usersDto.Status)
  return allUsers, err
}

func (userService *UserService) getMaps() map[string]interface{} {
  maps := make(map[string]interface{})
  maps["deleted_on"] = 0

  if userService.Status != "" {
    maps["status"] = userService.Status
  }

  return maps
}


func UserGetAll(db *gorm.DB, pageNum int, pageSize int, status string) ([]model.User, error) {
  userService := &UserService{}

  if status != "" {
    userService.Status = status
  }

  maps := userService.getMaps()
  allUsers, err := model.UserGetAll(db, pageSize, pageNum, maps)
  if err != nil {
    return nil, err
  }

  return allUsers, nil
}
