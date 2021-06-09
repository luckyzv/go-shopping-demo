package model

import "gorm.io/gorm"

type User struct {
  gorm.Model

  Name string  `json:"name" gorm:"size:16;not null,comment:用户姓名"`
  PassWord string `json:"password" gorm:"type:varchar(20);column:password;not null"`
  Age int `json:"age" gorm:"type:tinyint;unsigned;default:0"`
  Email string `json:"email" gorm:"type:varchar(30);unique"`
  Phone string  `json:"phone" gorm:"size:11;not null;uniqueIndex"`
  AvatarUrl string `json:"avatarUrl" gorm:"type:varchar(255)"`
  Status string `json:"status" gorm:"type:enum('published','pending','deleted');default:pending;comment:注册状态"`
}

//func (db *gorm.DB) ExistUserByPhone(phone string) (bool, error)  {
//
//}
