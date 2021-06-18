package router

import (
  "github.com/gin-gonic/gin"
	"shopping/src/controller"
)

func OrderRouters(e *gin.Engine)  {
  c := &controller.OrderController{}
  order := e.Group("/api/v1/orders")

  order.GET("/hello", c.CreateOrder)
}
