package main

import (
  "github.com/gin-gonic/gin"
  "github.com/swaggo/gin-swagger"
  "github.com/swaggo/gin-swagger/swaggerFiles"
  "net/http"
  "shopping/config"
  _ "shopping/docs"
  _ "shopping/engine"
  "shopping/middleware"
  "shopping/router"
)

// @Title Go-Shopping Swagger Example API
// @Version 1.0
// @Description This is a sample server Go-Shopping server.
// @Host localhost:3002
// @BasePath /api/v1/
func main() {
  r := gin.Default()

  //envName := os.Getenv("ENV")
  //if envName == "Staging" || envName == "Prod" {
  //  r.Use(middleware.LoggerToFile())
  //}

  // health check路由(consul使用)
  r.GET("/api/v1/Health", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "ok",
    })
    return
  })
  // swagger文档路由
  r.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "NOT_SHOWN_API_DOC"))

  // 加载业务路由
  r.Use(middleware.Logger())
  router.Init(r)

  viperConfig := config.GetServerConfig()
  port := viperConfig.Port
  if port != "" {
    r.Run(":" + port)
  } else {
    r.Run()
  }
}
