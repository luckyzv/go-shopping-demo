package router

import (
  "github.com/gin-gonic/gin"
  "shopping/controller"
)

func UserRouters(e *gin.Engine)  {
  user := e.Group("/api/v1/users")
  user.GET("/hello", controller.UserRegister)
  user.POST("/register", controller.UserRegister)
}
