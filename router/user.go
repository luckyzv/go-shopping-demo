package router

import (
  "github.com/gin-gonic/gin"
  "shopping/controller"
)

func UserRouters(e *gin.Engine)  {
  user := e.Group("/api/v1/users")
  user.POST("/register", controller.UserRegister) // 用户注册
  user.POST("/login", controller.UserLogin) // 用户登录
}
