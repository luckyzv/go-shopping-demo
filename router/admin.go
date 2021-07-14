package router

import (
  "github.com/gin-gonic/gin"
  "shopping/middleware"
  "shopping/src/controller"
)

func AdminRouters(e *gin.Engine)  {
  c := &controller.AdminController{}
  admin := e.Group("/api/v1/admins")
  admin.POST("/login", c.Login)

  admin.Use(middleware.Auth())
  admin.GET("/users", c.GetALlUsers)
}
