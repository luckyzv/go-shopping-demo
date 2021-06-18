package router

import (
  "github.com/gin-gonic/gin"
  "shopping/middleware"
  "shopping/src/controller"
)

func AdminRouters(e *gin.Engine)  {
  c := &controller.AdminController{}
  e.POST("/api/v1/admins/login", c.Login)
  admin := e.Group("/api/v1/admins")
  admin.Use(middleware.Auth())

  admin.GET("/hello", c.CreateProduct)
}
