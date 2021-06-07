package router

import (
  "github.com/gin-gonic/gin"
  "shopping/controller"
)

func OrderRouters(e *gin.Engine)  {
  order := e.Group("/api/v1/orders")
  order.GET("/hello", controller.UserRegister)
}
