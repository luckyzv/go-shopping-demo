package router

import (
  "github.com/gin-gonic/gin"
	"shopping/src/controller"
)

func ProductRouters(e *gin.Engine)  {
  c := &controller.ProductController{}
  product := e.Group("/api/v1/products")

  product.POST("", c.AddNewProduct)
  product.GET("/:productId", c.GetProduct)
}
