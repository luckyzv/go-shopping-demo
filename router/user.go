package router

import (
  "github.com/gin-gonic/gin"
  "shopping/src/controller"
)

func UserRouters(e *gin.Engine)  {
  c := &controller.UserController{}
  user := e.Group("/api/v1/users")

  user.POST("/register", c.UserRegister) // 用户注册
  user.POST("/login", c.UserLogin)       // 用户登录
}
