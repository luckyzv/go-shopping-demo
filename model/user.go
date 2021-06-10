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

func ExistUserByEmail(db *gorm.DB, email string) (bool, error)  {
  var user User
  err := db.Select("id").Where("email = ? AND deleted_at IS NULL", email).First(&user).Error
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

func GetUserById(db *gorm.DB, id string) (*User, error)  {
  var user User
  err := db.Where("id = ? AND deleted_at IS NULL", id).First(&user).Error
  if err != nil && err != gorm.ErrRecordNotFound {
    return nil, err
  }
  return &user, nil
}

func CreateUser(db *gorm.DB, data map[string]interface{}) error  {
  user := User{
    Name: data["name"].(string),
    PassWord: data["password"].(string),
    Age: data["age"].(int),
    Email: data["email"].(string),
    Phone: data["phone"].(string),
    Status: "pending",
  }
  if err := db.Create(&user).Error; err != nil {
    return err
  }
  return nil
}
