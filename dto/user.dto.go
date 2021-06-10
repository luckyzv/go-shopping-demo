package dto

import (
  "gorm.io/gorm"
  "shopping/model"
)

type UserDto struct {
  gorm.Model

  Name string  `json:"name"`
  PassWord string `json:"password"`
  Age int `json:"age"`
  Email string `json:"email"`
  Phone string  `json:"phone"`
  AvatarUrl string `json:"avatarUrl"`
  Status string `json:"status"`
}

type UserLoginDto struct {
  Email string `json:"email" binding:"required"`
  Password string `json:"password" binding:"required"`
}

func UserInfo(user model.User) UserDto {
  return UserDto{
    Name: user.Name,
    Phone: user.Phone,
    Age: user.Age,
    Email: user.Email,
    AvatarUrl: user.AvatarUrl,
    Status: user.Status,
  }
}

