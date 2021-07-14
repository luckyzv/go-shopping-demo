package router

import (
  "github.com/gin-gonic/gin"
  "shopping/middleware"
  "shopping/src/controller"
)

func UserRouters(e *gin.Engine)  {
  c := &controller.UserController{}
  user := e.Group("/api/v1/users")

  user.POST("/register", c.UserRegister) // 用户注册
  user.POST("/login", c.UserLogin)       // 用户登录

  user.Use(middleware.Auth())
  user.GET("/info", c.UserInfo) // 用户个人信息
  user.GET("", c.GetUsers)
}
