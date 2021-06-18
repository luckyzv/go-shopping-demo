package router

import (
  "github.com/gin-gonic/gin"
  "shopping/response"
)

type FuncRouter func(engine *gin.Engine)

var routers []FuncRouter
func Includes(funcRouters ...FuncRouter)  {
  routers = append(routers, funcRouters...)
}

func Init(r *gin.Engine) {
  Includes(OrderRouters, UserRouters, ProductRouters, AdminRouters)
  for _, funcRouter := range routers {
    funcRouter(r)
  }
  r.NoRoute(response.NotFound)
}
