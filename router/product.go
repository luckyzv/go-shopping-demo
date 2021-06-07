package router

import (
  "github.com/gin-gonic/gin"
  "shopping/controller"
)

func ProductRouters(e *gin.Engine)  {
  product := e.Group("/api/v1/products")
  product.GET("/hello", controller.UserRegister)
}
