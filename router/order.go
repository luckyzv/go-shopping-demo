package router

import (
	"github.com/gin-gonic/gin"
  "shopping/middleware"
  "shopping/src/order"
)

func OrderRouters(e *gin.Engine)  {
  c := &order.Controller{}
  orderRouter := e.Group("/api/v1/orders")

  orderRouter.Use(middleware.Auth())
  orderRouter.GET("/orderId", c.GetOrderId)
  orderRouter.POST("", c.AddNewOrder)
  orderRouter.POST("/sec-kill", c.AddNewSecKillOrder)
}
