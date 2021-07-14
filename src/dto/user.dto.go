package dto

import (
  "gorm.io/gorm"
  "shopping/model"
  "time"
)

type UserDto struct {
  Id uint `json:"id"`
  CreatedAt time.Time `json:"createdAt"`
  DeletedAt time.Time `json:"deletedAt"`
  Name string  `json:"name"`
  Age int `json:"age"`
  Email string `json:"email"`
  Phone string  `json:"phone"`
  AvatarUrl string `json:"avatarUrl"`
  Status string `json:"status"`
  Token string `json:"token"`
}

type UserLoginDto struct {
  Phone string `json:"phone" binding:"required"`
  Password string `json:"password" binding:"required"`
}

type GetAllUsersDto struct {
  Status string `json:"status"`
  PageSize int `json:"pageSize" binding:"required"`
  PageNum int `json:"pageNum" binding:"required"`
}

func UserLoginResponseDto(user model.User, token string) UserDto {
  return UserDto{
    Id: user.ID,
    CreatedAt: user.CreatedAt,
    Name: user.Name,
    Phone: user.Phone,
    Age: user.Age,
    Email: user.Email,
    AvatarUrl: user.AvatarUrl,
    Status: user.Status,
    Token: token,
  }
}

func ConvertUserDtoToModel(userDto UserDto) model.User  {
  return model.User{
    Model:     gorm.Model{
      ID: userDto.Id,
      CreatedAt: userDto.CreatedAt,
    },
    Name:      userDto.Name,
    Age:       userDto.Age,
    Email:     userDto.Email,
    Phone:     userDto.Phone,
    AvatarUrl: userDto.AvatarUrl,
    Status:    userDto.Status,
  }
}

func ConvertModelUserToDto(user model.User) UserDto  {
  return UserDto{
    Id: user.ID,
    CreatedAt: user.CreatedAt,
    Name:      user.Name,
    Age:       user.Age,
    Email:     user.Email,
    Phone:     user.Phone,
    AvatarUrl: user.AvatarUrl,
    Status:    user.Status,
  }
}
