package model

import (
  "gorm.io/gorm"
)

type User struct {
  gorm.Model

  Name string  `json:"name" gorm:"size:16;not null,comment:用户姓名"`
  PassWord string `json:"password" gorm:"type:varchar(100);column:password;not null"`
  Age int `json:"age" gorm:"type:tinyint;unsigned;default:0"`
  Email string `json:"email" gorm:"type:varchar(30);not null;default:''"`
  Phone string  `json:"phone" gorm:"size:11;not null;uniqueIndex"`
  AvatarUrl string `json:"avatarUrl" gorm:"type:varchar(255);not null;default:''"`
  Status string `json:"status" gorm:"type:enum('published','pending','deleted');default:pending;comment:注册状态"`
}

func ExistUserByPhone(db *gorm.DB, phone string) (bool, error)  {
  var user User
  err := db.Select("id").Where("phone = ? AND deleted_at IS NULL", phone).First(&user).Error
  if err != nil && err != gorm.ErrRecordNotFound {
    return false, err
  }
  if user.ID > 0 {
    return true, nil
  }
  return false, nil
}

func GetUsers()  {

}

func GetUserById(db *gorm.DB, id uint) (*User, error)  {
  var user User
  err := db.Where("id = ? AND deleted_at IS NULL", id).First(&user).Error
  if err != nil && err != gorm.ErrRecordNotFound {
    return nil, err
  }
  return &user, nil
}

func GetUserByPhone(db *gorm.DB, phone string) (bool, *User) {
  var user User
  err := db.Where("phone = ? AND deleted_at IS NULL", phone).First(&user).Error
  if err != nil && err == gorm.ErrRecordNotFound {
    return false, nil
  }
  return true, &user
}

func CreateUser(db *gorm.DB, user User) error  {
  if err := db.Create(&user).Error; err != nil {
    return err
  }
  return nil
}
